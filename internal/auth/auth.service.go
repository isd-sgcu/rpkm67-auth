package auth

import (
	"context"

	"github.com/isd-sgcu/rpkm67-auth/internal/user"
	proto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/auth/v1"
	userProto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/user/v1"
	"go.uber.org/zap"
)

type Service interface {
	proto.AuthServiceServer
}

type serviceImpl struct {
	proto.UnimplementedAuthServiceServer
	userSvc user.Service
	log     *zap.Logger
}

func NewService(log *zap.Logger) Service {
	return &serviceImpl{
		log: log,
	}
}

func (s *serviceImpl) Validate(_ context.Context, in *proto.ValidateRequest) (res *proto.ValidateResponse, err error) {
	return nil, nil
}

func (s *serviceImpl) RefreshToken(_ context.Context, in *proto.RefreshTokenRequest) (res *proto.RefreshTokenResponse, err error) {
	return nil, nil
}

func (s *serviceImpl) SignUp(_ context.Context, in *proto.SignUpRequest) (res *proto.SignUpResponse, err error) {
	createUser := &userProto.CreateUserRequest{
		Email:     in.Email,
		Password:  in.Password,
		Firstname: in.Firstname,
		Lastname:  in.Lastname,
		Role:      in.Role,
	}

	userRes, err := s.userSvc.Create(context.Background(), createUser)
	if err != nil {
		return nil, err
	}

	return &proto.SignUpResponse{
		Credential: nil,
	}, nil
}

func (s *serviceImpl) SignIn(_ context.Context, in *proto.SignInRequest) (res *proto.SignInResponse, err error) {
	// user := &model.User{}
	user, err := s.userSvc.FindByEmail(context.Background(), &userProto.FindByEmailRequest{})
	if err != nil {
		return nil, err
	}

	// return nil, nil
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
