package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// Id       int64  `json:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"`
	Tel      string `json:"tel"`
	Role     int    `json:"role"`
}
