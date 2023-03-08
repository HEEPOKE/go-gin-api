package user

import (
	"Backend/go-api/config"
	"Backend/go-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var users []model.User
	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "success",
			"users":   users,
		})

	}
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	var user []model.User
	config.DB.First(&user, id)
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "success",
		"users":   user,
	})
}

func Profile(c *gin.Context) {
	userId := c.MustGet("userId").(float64)
	var user []model.User
	config.DB.First(&user, userId)
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "success",
		"users":   user,
	})
}
