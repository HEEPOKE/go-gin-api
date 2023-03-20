package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ComparePasswords(c *gin.Context, hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":      "error",
			"message":     "Incorrect password",
			"description": "password",
		})
		return false
	}
	return true
}
