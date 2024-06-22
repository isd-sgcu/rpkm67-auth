package auth

import (
	"context"

	"github.com/isd-sgcu/rpkm67-auth/constant"
	"github.com/isd-sgcu/rpkm67-auth/internal/token"
	"github.com/isd-sgcu/rpkm67-auth/internal/user"
	proto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/auth/v1"
	userProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/user/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service interface {
	proto.AuthServiceServer
}

type serviceImpl struct {
	proto.UnimplementedAuthServiceServer
	userSvc  user.Service
	tokenSvc token.Service
	utils    AuthUtils
	bcrypt   BcryptUtils
	log      *zap.Logger
}

func NewService(userSvc user.Service, tokenSvc token.Service, utils AuthUtils, bcrypt BcryptUtils, log *zap.Logger) Service {
	return &serviceImpl{
		userSvc:  userSvc,
		tokenSvc: tokenSvc,
		utils:    utils,
		bcrypt:   bcrypt,
		log:      log,
	}
}

func (s *serviceImpl) Validate(_ context.Context, in *proto.ValidateRequest) (res *proto.ValidateResponse, err error) {
	return nil, nil
}

func (s *serviceImpl) RefreshToken(_ context.Context, in *proto.RefreshTokenRequest) (res *proto.RefreshTokenResponse, err error) {
	return nil, nil
}

func (s *serviceImpl) SignUp(_ context.Context, in *proto.SignUpRequest) (res *proto.SignUpResponse, err error) {
	role := "user"
	if s.utils.IsStudentIdInMap(in.Email) {
		role = "staff"
	}

	hashedPassword, err := s.bcrypt.GenerateHashedPassword(in.Password)
	if err != nil {
		s.log.Named("SignUp").Error("GenerateHashedPassword: ", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	createUser := &userProto.CreateUserRequest{
		Email:     in.Email,
		Password:  hashedPassword,
		Firstname: in.Firstname,
		Lastname:  in.Lastname,
		Role:      role,
	}

	userRes, err := s.userSvc.Create(context.Background(), createUser)
	if err != nil {
		s.log.Named("SignUp").Error("Create: ", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.SignUpResponse{
		Id:        userRes.User.Id,
		Email:     userRes.User.Email,
		Firstname: userRes.User.Firstname,
		Lastname:  userRes.User.Lastname,
	}, nil
}

func (s *serviceImpl) SignIn(_ context.Context, in *proto.SignInRequest) (res *proto.SignInResponse, err error) {
	user, err := s.userSvc.FindByEmail(context.Background(), &userProto.FindByEmailRequest{Email: in.Email})
	if err != nil {
		s.log.Named("SignIn").Error("FindByEmail: ", zap.Error(err))
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	err = s.bcrypt.CompareHashedPassword(user.User.Password, in.Password)
	if err != nil {
		s.log.Named("SignIn").Error("CompareHashedPassword: ", zap.Error(err))
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	credentials, err := s.tokenSvc.GetCredentials(user.User.Id, constant.Role(user.User.Role))
	if err != nil {
		s.log.Named("SignIn").Error("GetCredentials: ", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.SignInResponse{
		Credential: &proto.Credential{
			AccessToken:  credentials.AccessToken,
			RefreshToken: credentials.RefreshToken,
			ExpiresIn:    int32(s.tokenSvc.GetConfig().AccessTTL),
		},
	}, nil
}

func (s *serviceImpl) SignOut(_ context.Context, in *proto.SignOutRequest) (res *proto.SignOutResponse, err error) {
	return nil, nil
}

func (s *serviceImpl) ForgotPassword(_ context.Context, in *proto.ForgotPasswordRequest) (res *proto.ForgotPasswordResponse, err error) {
	return nil, nil
}

func (s *serviceImpl) ResetPassword(_ context.Context, in *proto.ResetPasswordRequest) (res *proto.ResetPasswordResponse, err error) {
	return nil, nil
}
