package model

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	UserName string
	Password string
}
