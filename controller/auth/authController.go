package auth

import (
	"Backend/go-api/common"
	"Backend/go-api/config"
	"Backend/go-api/model"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Register(c *gin.Context) {
	var json model.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userExists, _ := common.CheckUserExistence(json.Username); userExists {
		c.JSON(http.StatusOK, gin.H{
			"message": "user Exist",
			"status":  "error",
		})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to encrypt password",
			"status":  "error",
		})
		return
	}

	user := model.User{
		Username: json.Username,
		Password: string(encryptedPassword),
		Email:    json.Email,
		Tel:      json.Tel,
		Role:     model.Viewer,
	}
	config.DB.Create(&user)

	if user.ID <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Register Success",
		"userId":  user.ID,
	})
}

func Login(c *gin.Context) {
	var json model.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userExist, err := common.GetUser(json.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve user",
			"status":  "error",
		})
		return
	}

	if userExist.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "User does not exist",
		})
		return
	}

	if err := common.ComparePasswords(userExist.Password, json.Password); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "Incorrect password",
		})
		return
	}

	userID := int(userExist.ID)
	token, err := common.GenerateToken(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to generate token",
			"status":  "error",
		})
		return
	}

	parsedToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "Failed to parse token",
		})
		return
	}

	claims, ok := parsedToken.Claims.(*jwt.StandardClaims)
	if !ok || !parsedToken.Valid {
		c.JSON(http.StatusOK, gin.H{
			"status":  "error",
			"message": "Failed to retrieve claims from token",
		})
		return
	}

	payload := map[string]interface{}{
		"userId": userExist.ID,
		"role":   userExist.Role,
		"token":  token,
		"exp":    claims.ExpiresAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Login Success",
		"payload": payload,
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout Success",
	})
}
