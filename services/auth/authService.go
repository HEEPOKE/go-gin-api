package auth

import (
	"Backend/go-api/common"
	"Backend/go-api/config"
	"Backend/go-api/model"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	CheckUserExistence(username string) (bool, error)
	RegisterUser(user *model.User) error
}

type AuthServiceImpl struct{}

func NewAuthService() AuthService {
	return &AuthServiceImpl{}
}

func (s *AuthServiceImpl) RegisterUser(user *model.User) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}

	user.Password = string(encryptedPassword)
	result := config.DB.Create(user)
	return result.Error
}

func (s *AuthServiceImpl) CheckUserExistence(username string) (bool, error) {
	return common.CheckUserExistence(username)
}
