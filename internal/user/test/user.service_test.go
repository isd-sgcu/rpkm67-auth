package test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type UserServiceTest struct {
	suite.Suite
	// controller *gomock.Controller
	// logger     *zap.Logger
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(UserServiceTest))
}

func (t *UserServiceTest) SetupTest() {}

func (t *UserServiceTest) TestSignUpSuccess() {

}
