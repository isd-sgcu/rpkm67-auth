package auth

import (
	"context"

	proto "github.com/isd-sgcu/rpkm67-go-proto/rpkm67/auth/auth/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Service interface {
	proto.AuthServiceClient
}

type serviceImpl struct {
	proto.UnimplementedAuthServiceServer
	// repo file.Repository
	log *zap.Logger
}

func NewService(log *zap.Logger) Service {
	return &serviceImpl{
		log: log,
	}
}

func (s *serviceImpl) Validate(_ context.Context, in *proto.ValidateRequest, _ ...grpc.CallOption) (*proto.ValidateResponse, error) {
	return nil, nil
}

func (s *serviceImpl) RefreshToken(_ context.Context, in *proto.RefreshTokenRequest, _ ...grpc.CallOption) (*proto.RefreshTokenResponse, error) {
	return nil, nil
}

func (s *serviceImpl) SignUp(_ context.Context, in *proto.SignUpRequest, _ ...grpc.CallOption) (*proto.SignUpResponse, error) {
	return nil, nil
}

func (s *serviceImpl) SignIn(_ context.Context, in *proto.SignInRequest, _ ...grpc.CallOption) (*proto.SignInResponse, error) {
	return nil, nil
}

func (s *serviceImpl) SignOut(_ context.Context, in *proto.SignOutRequest, _ ...grpc.CallOption) (*proto.SignOutResponse, error) {
	return nil, nil
}

func (s *serviceImpl) ForgotPassword(_ context.Context, in *proto.ForgotPasswordRequest, _ ...grpc.CallOption) (*proto.ForgotPasswordResponse, error) {
	return nil, nil
}

func (s *serviceImpl) ResetPassword(_ context.Context, in *proto.ResetPasswordRequest, _ ...grpc.CallOption) (*proto.ResetPasswordResponse, error) {
	return nil, nil
}
