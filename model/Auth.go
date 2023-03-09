package model

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	UsernameOrEmail string `json:"username_or_email" binding:"required"`
	Password        string `json:"password" binding:"required"`
}
