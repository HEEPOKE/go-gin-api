package common

import (
	"Backend/go-api/config"
	"Backend/go-api/model"
	"errors"

	"gorm.io/gorm"
)

func GetUserByUsernameOrEmail(usernameOrEmail string) (*model.User, error) {
	var user model.User
	err := config.DB.Where("username = ?", usernameOrEmail).Or("email = ?", usernameOrEmail).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetUser(username string) (model.User, error) {
	var userExist model.User
	err := config.DB.Where("username = ?", username).First(&userExist).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.User{}, nil
		}
		return model.User{}, err
	}
	return userExist, nil
}
