package model

import "gorm.io/gorm"

type UserRole string

const (
	Admin  UserRole = "admin"
	Editor UserRole = "editor"
	Viewer UserRole = "viewer"
)

type User struct {
	gorm.Model
	Username string   `json:"username"`
	Password string   `json:"password"`
	Email    string   `json:"email"`
	Tel      string   `json:"tel"`
	Role     UserRole `json:"role"`
}
