package controller

import (
	"Backend/go-api/model"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func register() {
	dsn := "root:root@tcp(127.0.0.1:3306)/Shirtgo?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	r := gin.Default()
	r.Use(cors.Default())

	db.AutoMigrate(&model.User{})

	r.POST("/register", func(c *gin.Context) {
		var json model.User
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var userExist model.User
		db.Where("username = ?", json.Username).First(&userExist)
		if userExist.ID > 0 {
			c.JSON(http.StatusOK, gin.H{
				"message": "user Exist",
				"status":  "error",
			})
			return
		}

		encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
		user := model.User{Username: json.Username, Password: string(encryptedPassword), Email: json.Email}
		db.Create(&user)
		if user.ID > 0 {
			c.JSON(http.StatusOK, gin.H{
				"userId":  user.ID,
				"message": "success",
				"status":  "ok",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "fail",
				"status":  "error",
			})
		}
	})
}
