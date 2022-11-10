package user

import (
	"Backend/go-api/config"
	"Backend/go-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var users []model.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "success",
	})
}
