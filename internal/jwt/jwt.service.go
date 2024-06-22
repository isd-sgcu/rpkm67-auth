package jwt

import (
	"fmt"
	"time"

	_jwt "github.com/golang-jwt/jwt/v4"
	"github.com/isd-sgcu/rpkm67-auth/config"
	"github.com/isd-sgcu/rpkm67-auth/constant"
	"github.com/isd-sgcu/rpkm67-auth/internal/dto"
	"github.com/pkg/errors"
)

type Service interface {
	SignAuth(userId string, role constant.Role, authSessionId string) (string, error)
	VerifyAuth(token string) (*_jwt.Token, error)
	GetConfig() *config.JwtConfig
}

type serviceImpl struct {
	config   config.JwtConfig
	strategy JwtStrategy
	jwtUtils JwtUtils
}

func NewService(config config.JwtConfig, strategy JwtStrategy, jwtUtils JwtUtils) Service {
	return &serviceImpl{config: config, strategy: strategy, jwtUtils: jwtUtils}
}

func (s *serviceImpl) SignAuth(userId string, role constant.Role, authSessionId string) (string, error) {
	payloads := dto.AuthPayload{
		RegisteredClaims: _jwt.RegisteredClaims{
			Issuer:    s.config.Issuer,
			ExpiresAt: s.jwtUtils.GetNumericDate(time.Now().Add(time.Second * time.Duration(s.config.AccessTTL))),
			IssuedAt:  s.jwtUtils.GetNumericDate(time.Now()),
		},
		UserID:        userId,
		AuthSessionID: authSessionId,
	}

	token := s.jwtUtils.GenerateJwtToken(_jwt.SigningMethodHS256, payloads)

	tokenStr, err := s.jwtUtils.SignedTokenString(token, s.config.Secret)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error while signing the token due to: %s", err.Error()))
	}

	return tokenStr, nil
}

func (s *serviceImpl) VerifyAuth(token string) (*_jwt.Token, error) {
	return s.jwtUtils.ParseToken(token, s.strategy.AuthDecode)
}

func (s *serviceImpl) GetConfig() *config.JwtConfig {
	return &s.config
}
