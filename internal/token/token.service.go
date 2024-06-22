package token

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/isd-sgcu/rpkm67-auth/config"
	"github.com/isd-sgcu/rpkm67-auth/constant"
	"github.com/isd-sgcu/rpkm67-auth/internal/cache"
	"github.com/isd-sgcu/rpkm67-auth/internal/dto"
	"github.com/isd-sgcu/rpkm67-auth/internal/jwt"
)

type Service interface {
	GetCredentials(userId string, role constant.Role) (*dto.Credentials, error)
	CreateCredentials(userId string, role constant.Role) (*dto.Credentials, error)
	RefreshToken(refreshToken string) (*dto.Credentials, error)
	GetConfig() *config.JwtConfig
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

func (s *serviceImpl) GetCredentials(userId string, role constant.Role) (*dto.Credentials, error) {
	credentials := &dto.Credentials{}
	err := s.cache.GetValue(sessionKey(userId), credentials)
	if err != nil {
		return nil, err
	}

	if (credentials == &dto.Credentials{}) { // no session found
		credentials, err = s.CreateCredentials(userId, role)
		if err != nil {
			return nil, err
		}
	}

	_, err = s.jwtService.ValidateToken(credentials.AccessToken)
	if err != nil { // still have refreshToken but accessToken is expired
		err := s.cache.DeleteValue(sessionKey(userId))
		if err != nil {
			return nil, err
		}

		accessToken, err := s.jwtService.CreateToken(userId, role)
		if err != nil {
			return nil, err
		}

		newCredentials := &dto.Credentials{
			AccessToken:  accessToken,
			RefreshToken: credentials.RefreshToken,
		}

		err = s.cache.SetValue(sessionKey(userId), newCredentials, s.jwtService.GetConfig().AccessTTL)
		if err != nil {
			return nil, err
		}

		return newCredentials, nil
	}

	return credentials, nil
}

func (s *serviceImpl) CreateCredentials(userId string, role constant.Role) (*dto.Credentials, error) {
	accessToken, err := s.jwtService.CreateToken(userId, role)
	if err != nil {
		return nil, err
	}

	refreshToken := createRefreshToken()

	err = s.cache.SetValue(refreshKey(refreshToken), &dto.RefreshTokenCache{
		UserID: userId,
		Role:   role,
	}, s.jwtService.GetConfig().RefreshTTL)
	if err != nil {
		return nil, err
	}

	credentials := &dto.Credentials{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	err = s.cache.SetValue(sessionKey(userId), credentials, s.jwtService.GetConfig().AccessTTL)
	if err != nil {
		return nil, err
	}

	return credentials, nil
}

func (s *serviceImpl) RefreshToken(refreshToken string) (*dto.Credentials, error) {
	refreshCache := &dto.RefreshTokenCache{}
	err := s.cache.GetValue(refreshKey(refreshToken), refreshCache)
	if err != nil {
		return nil, err
	} else if (refreshCache == &dto.RefreshTokenCache{}) {
		return nil, fmt.Errorf("refresh token not found")
	}

	err = s.cache.DeleteValue(refreshKey(refreshToken))
	if err != nil {
		return nil, err
	}

	err = s.cache.DeleteValue(sessionKey(refreshCache.UserID))
	if err != nil {
		return nil, err
	}

	credentials, err := s.CreateCredentials(refreshCache.UserID, refreshCache.Role)
	if err != nil {
		return nil, err
	}

	return credentials, nil
}

func (s *serviceImpl) GetConfig() *config.JwtConfig {
	return s.jwtService.GetConfig()
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
