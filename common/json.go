package common

import "github.com/gin-gonic/gin"

func BindJSON(c *gin.Context, obj interface{}) error {
	return c.ShouldBindJSON(obj)
}
