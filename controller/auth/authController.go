package auth

import (
	"Backend/go-api/common"
	"Backend/go-api/config"
	"Backend/go-api/model"
	"Backend/go-api/services/auth"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	UserService auth.UserService
}

const SecretKey = "secret"

func (a *Auth) Register(c *gin.Context) {
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

	if err := config.DB.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
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
	var json model.Auth
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := common.GetUserByUsernameOrEmail(json.UsernameOrEmail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":     "Failed to retrieve user",
			"status":      "error",
			"description": "",
		})
		return
	}

	if user == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":      "error",
			"message":     "Invalid username or email",
			"description": "username or email",
		})
		return
	}

	if err := common.ComparePasswords(user.Password, json.Password); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":      "error",
			"message":     "Incorrect password",
			"description": "password",
		})
		return
	}

	token, err := common.GenerateToken(int(user.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":     "Failed to generate token",
			"status":      "error",
			"description": "",
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
			"status":      "error",
			"message":     "Failed to parse token",
			"description": "",
		})
		return
	}

	claims, ok := parsedToken.Claims.(*jwt.StandardClaims)
	if !ok || !parsedToken.Valid {
		c.JSON(http.StatusOK, gin.H{
			"status":      "error",
			"message":     "Failed to retrieve claims from token",
			"description": "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Login Success",
		"payload": gin.H{
			"userId":   strconv.FormatUint(uint64(user.ID), 10),
			"username": user.Username,
			"email":    user.Email,
			"tel":      user.Tel,
			"role":     user.Role,
			"token":    token,
			"exp":      claims.ExpiresAt,
		},
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout Success",
	})
}
