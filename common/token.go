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

func ParseToken(tokenString string, secretKey []byte) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return secretKey, nil
	})
}
