package product

import (
	"Backend/go-api/config"
	"Backend/go-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadProduct(c *gin.Context) {
	var product []model.Product
	config.DB.First(&product)
	c.JSON(http.StatusOK, product)
}

func Create(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	result := config.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{
		"Create": result.RowsAffected,
	})
}

func Edit(c *gin.Context) {
	id := c.Param("id")
	var product model.Product
	var productUpdate model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	config.DB.First(&productUpdate, id)
	result := config.DB.Save(&productUpdate)
	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	config.DB.First(&product, id)
	result := config.DB.Delete(&product)
	c.JSON(http.StatusOK, result)
}
