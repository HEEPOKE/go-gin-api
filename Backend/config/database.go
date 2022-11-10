package config

import (
	"Backend/go-api/model"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Database() {
	dsn := os.Getenv("MYSQL_DB")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Product{})
}
