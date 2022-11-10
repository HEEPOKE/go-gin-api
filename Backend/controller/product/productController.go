package product

import (
	"Backend/go-api/config"
	"Backend/go-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	result := config.DB.Create(&product)
	c.JSON(http.StatusOK, result)
}

func Edit(c *gin.Context) {
	id := c.Param("id")
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	result := config.DB.First(&product, id)
	c.JSON(http.StatusOK, result)
}
