package model

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
