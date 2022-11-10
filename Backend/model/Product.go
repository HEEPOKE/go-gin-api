package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string `json:"username" binding:"required"`
	Color    string `json:"email" binding:"required"`
	Category string `json:"pssword" binding:"required"`
	Price    uint   `json:"price" binding:"required"`
}
