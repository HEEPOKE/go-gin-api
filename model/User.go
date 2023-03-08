package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Tel      string `json:"tel"`
	Role     int    `json:"role"`
}
