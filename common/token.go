package common

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(userID int) (string, error) {
	hmacSampleSecret := []byte(os.Getenv("JWT_SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userID,
		"exp":    time.Now().Add(time.Minute * 1).Unix(),
	})
	return token.SignedString(hmacSampleSecret)
}
