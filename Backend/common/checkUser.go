package common

import (
	"Backend/go-api/config"
	"Backend/go-api/model"

	"gorm.io/gorm"
)

func CheckUserExistence(username string) (bool, error) {
	var userExist model.User
	err := config.DB.Where("username = ?", username).First(&userExist).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
