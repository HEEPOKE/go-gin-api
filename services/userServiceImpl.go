package services

import (
	"Backend/go-api/common"
	"Backend/go-api/config"
	"Backend/go-api/model"
)

type UserServiceImpl struct{}

func (s *UserServiceImpl) CheckUserExistence(username string) (bool, error) {
	return common.CheckUserExistence(username)
}

func (s *UserServiceImpl) CreateUser(user *model.User) error {
	config.DB.Create(user)
	return nil
}
