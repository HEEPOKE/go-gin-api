package model

import "gorm.io/gorm"

type RestErr struct {
	gorm.Model
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}
