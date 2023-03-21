package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func parseToken(tokenString string, secretKey []byte) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
}

func ValidationUsers() gin.HandlerFunc {
	hmacSampleSecret := []byte(os.Getenv("JWT_SECRET_KEY"))

	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		tokenString := strings.Replace(header, "Bearer ", "", 1)

		token, err := parseToken(tokenString, hmacSampleSecret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status":  "forbidden",
				"message": err.Error(),
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userId", claims["userId"])
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status":  "forbidden",
				"message": errors.New("invalid token").Error(),
			})
			return
		}
		c.Next()
	}
}
