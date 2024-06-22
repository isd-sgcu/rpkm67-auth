package token

import (
	"github.com/isd-sgcu/rpkm67-auth/internal/cache"
	"github.com/isd-sgcu/rpkm67-auth/internal/jwt"
)

type Service interface{}

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
