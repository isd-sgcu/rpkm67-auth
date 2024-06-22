package token

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/isd-sgcu/rpkm67-auth/internal/cache"
	"github.com/isd-sgcu/rpkm67-auth/internal/dto"
	"github.com/isd-sgcu/rpkm67-auth/internal/jwt"
	"github.com/isd-sgcu/rpkm67-auth/internal/model"
)

type Service interface {
	GetCredentials(user *model.User) (*dto.Credentials, error)
	CreateCredentials(user *model.User) (*dto.Credentials, error)
	RefreshToken(refreshToken string) (*dto.Credentials, error)
}

type serviceImpl struct {
	jwtService jwt.Service
	cache      cache.Repository
	tokenUtils TokenUtils
}

func NewService(jwtService jwt.Service, cache cache.Repository, tokenUtils TokenUtils) Service {
	return &serviceImpl{
		jwtService: jwtService,
		cache:      cache,
		tokenUtils: tokenUtils,
	}
}

func (s *serviceImpl) GetCredentials(user *model.User) (*dto.Credentials, error) {
	credentials := &dto.Credentials{}
	err := s.cache.GetValue(sessionKey(user.ID.String()), credentials)
	if err != nil {
		return nil, err
	}

	if (credentials == &dto.Credentials{}) { // no session found
		credentials, err = s.CreateCredentials(user)
		if err != nil {
			return nil, err
		}
	}

	_, err = s.jwtService.ValidateToken(credentials.AccessToken)
	if err != nil { // still have refreshToken but accessToken is expired
		err := s.cache.DeleteValue(sessionKey(user.ID.String()))
		if err != nil {
			return nil, err
		}

		accessToken, err := s.jwtService.CreateToken(user)
		if err != nil {
			return nil, err
		}

		newCredentials := &dto.Credentials{
			AccessToken:  accessToken,
			RefreshToken: credentials.RefreshToken,
		}

		err = s.cache.SetValue(sessionKey(user.ID.String()), newCredentials, s.jwtService.GetConfig().AccessTTL)
		if err != nil {
			return nil, err
		}

		return newCredentials, nil
	}

	return credentials, nil
}

func (s *serviceImpl) CreateCredentials(user *model.User) (*dto.Credentials, error) {
	accessToken, err := s.jwtService.CreateToken(user)
	if err != nil {
		return nil, err
	}

	refreshToken := createRefreshToken()

	err = s.cache.SetValue(refreshKey(refreshToken), user, s.jwtService.GetConfig().RefreshTTL)
	if err != nil {
		return nil, err
	}

	credentials := &dto.Credentials{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	err = s.cache.SetValue(sessionKey(user.ID.String()), credentials, s.jwtService.GetConfig().AccessTTL)
	if err != nil {
		return nil, err
	}

	return credentials, nil
}

func (s *serviceImpl) RefreshToken(refreshToken string) (*dto.Credentials, error) {
	user := &model.User{}
	err := s.cache.GetValue(refreshKey(refreshToken), user)
	if err != nil {
		return nil, err
	} else if (user == &model.User{}) {
		return nil, fmt.Errorf("refresh token not found")
	}

	err = s.cache.DeleteValue(refreshKey(refreshToken))
	if err != nil {
		return nil, err
	}

	err = s.cache.DeleteValue(sessionKey(user.ID.String()))
	if err != nil {
		return nil, err
	}

	credentials, err := s.CreateCredentials(user)
	if err != nil {
		return nil, err
	}

	return credentials, nil
}

func createRefreshToken() string {
	return uuid.New().String()
}

func refreshKey(refreshToken string) string {
	return fmt.Sprintf("refresh:%s", refreshToken)
}

func sessionKey(userId string) string {
	return fmt.Sprintf("session:%s", userId)
}
