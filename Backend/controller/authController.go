package controller

import (
	"Backend/go-api/config"
	"Backend/go-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var json model.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var userExist model.User
	config.DB.Where("username = ?", json.Username).First(&userExist)
	if userExist.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "user Exist",
			"status":  "error",
		})
		return
	}
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
	user := model.User{Username: json.Username, Password: string(encryptedPassword), Email: json.Email}
	config.DB.Create(&user)
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
}

func Login(c *gin.Context) {
	var json model.Auth
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var userExist model.Auth
	config.DB.Where("username = ?", json.Username).First(&userExist)
	if userExist.ID > 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "user Exist",
			"status":  "error",
		})
		return
	}
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
	user := model.Auth{Username: json.Username, Password: string(encryptedPassword)}
	config.DB.Create(&user)
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
}
