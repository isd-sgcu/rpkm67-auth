package user

import (
	"context"
	"errors"

	proto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/user/v1"
	"github.com/isd-sgcu/rpkm67-model/constant"
	"github.com/isd-sgcu/rpkm67-model/model"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type Service interface {
	proto.UserServiceServer
}

type serviceImpl struct {
	proto.UnimplementedUserServiceServer
	repo Repository
	log  *zap.Logger
}

func NewService(repo Repository, log *zap.Logger) proto.UserServiceServer {
	return &serviceImpl{
		repo: repo,
		log:  log,
	}
}

func (s *serviceImpl) Create(_ context.Context, req *proto.CreateUserRequest) (res *proto.CreateUserResponse, err error) {
	createUser := &model.User{
		Email: req.Email,
		Role:  constant.Role(req.Role),
	}

	err = s.repo.Create(createUser)
	if err != nil {
		s.log.Named("Create").Error("Create: ", zap.Error(err))
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, status.Error(codes.AlreadyExists, "duplicate email")
		}
		return nil, err
	}

	return &proto.CreateUserResponse{
		User: ModelToProto(createUser),
	}, nil
}

func (s *serviceImpl) FindOne(_ context.Context, req *proto.FindOneUserRequest) (res *proto.FindOneUserResponse, err error) {
	user := &model.User{}

	err = s.repo.FindOne(req.Id, user)
	if err != nil {
		s.log.Named("FindOne").Error("FindOne: ", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.FindOneUserResponse{
		User: ModelToProto(user),
	}, nil
}

func (s *serviceImpl) FindByEmail(_ context.Context, req *proto.FindByEmailRequest) (res *proto.FindByEmailResponse, err error) {
	user := &model.User{}

	err = s.repo.FindByEmail(req.Email, user)
	if err != nil {
		s.log.Named("FindByEmail").Error("FindByEmail: ", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.FindByEmailResponse{
		User: ModelToProto(user),
	}, nil
}

func (s *serviceImpl) Update(_ context.Context, req *proto.UpdateUserRequest) (res *proto.UpdateUserResponse, err error) {
	// pun devops
	return nil, nil
}

func ModelToProto(in *model.User) *proto.User {
	return &proto.User{
		Id:        in.ID.String(),
		Email:     in.Email,
		Firstname: in.Firstname,
		Lastname:  in.Lastname,
		Role:      in.Role.String(),
	}
}

func ProtoToModel(in *proto.User) *model.User {
	return &model.User{
		Email:     in.Email,
		Firstname: in.Firstname,
		Lastname:  in.Lastname,
		Role:      constant.Role(in.Role),
	}
}
