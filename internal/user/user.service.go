package user

import (
	"context"
	"errors"

	"github.com/isd-sgcu/rpkm67-auth/internal/constant"
	"github.com/isd-sgcu/rpkm67-auth/internal/model"
	proto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/user/v1"
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
	repo  Repository
	utils UserUtils
	log   *zap.Logger
}

func NewService(repo Repository, utils UserUtils, log *zap.Logger) proto.UserServiceServer {
	return &serviceImpl{
		repo:  repo,
		utils: utils,
		log:   log,
	}
}

func (s *serviceImpl) Create(_ context.Context, req *proto.CreateUserRequest) (res *proto.CreateUserResponse, err error) {
	hashPassword, err := s.utils.GenerateHashedPassword(req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, constant.InternalServerErrorMessage)
	}

	createUser := &model.User{
		StudentId: req.StudentId,
		Password:  hashPassword,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Tel:       req.Tel,
		Role:      constant.USER,
	}

	err = s.repo.Create(createUser)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, status.Error(codes.AlreadyExists, constant.DuplicateEmailErrorMessage)
		}
		return nil, status.Error(codes.Internal, constant.InternalServerErrorMessage)
	}

	return &proto.CreateUserResponse{
		User: ModelToProto(createUser),
	}, nil
}

func (s *serviceImpl) FindOne(_ context.Context, req *proto.FindOneUserRequest) (res *proto.FindOneUserResponse, err error) {
	user := &model.User{}

	err = s.repo.FindOne(req.Id, user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, constant.UserNotFoundErrorMessage)
		}
		return nil, status.Error(codes.Internal, constant.InternalServerErrorMessage+" "+err.Error())
	}
	return &proto.FindOneUserResponse{
		User: ModelToProto(user),
	}, nil
}

func ModelToProto(in *model.User) *proto.User {
	return &proto.User{
		Id:        in.ID.String(),
		StudentId: in.StudentId,
		Firstname: in.Firstname,
		Lastname:  in.Lastname,
		Tel:       in.Tel,
		Role:      in.Role.String(),
	}
}
