package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string `json:"name"`
	Category string `json:"category"`
	Color    string `json:"color"`
	Price    uint   `json:"price"`
}
