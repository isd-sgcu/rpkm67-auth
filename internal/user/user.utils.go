package user

import "golang.org/x/crypto/bcrypt"

type UserUtils interface {
	GenerateHashedPassword(password string) (string, error)
	CompareHashedPassword(hashedPassword string, plainPassword string) error
}

type userUtilsImpl struct{}

func NewUserUtils() UserUtils {
	return &userUtilsImpl{}
}

func (u *userUtilsImpl) GenerateHashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func (u *userUtilsImpl) CompareHashedPassword(hashedPassword string, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}
