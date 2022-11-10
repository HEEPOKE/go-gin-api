package product

import (
	"Backend/go-api/config"
	"Backend/go-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	userId := c.MustGet("userId").(float64)
	var user []model.User
	config.DB.First(&user, userId)
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "success",
		"users":   user,
	})
}
