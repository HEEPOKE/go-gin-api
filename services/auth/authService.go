package auth

import (
	"Backend/go-api/common"
	"Backend/go-api/config"
	"Backend/go-api/model"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	CheckUserExistence(username string) (bool, error)
	RegisterUser(user *model.User) error
	RespondWithToken(c *gin.Context, user *model.User)
	GetUserByUsernameOrEmail(c *gin.Context, usernameOrEmail string) (*model.User, error)
}

type AuthServiceImpl struct{}

func NewAuthService() AuthService {
	return &AuthServiceImpl{}
}

func (s *AuthServiceImpl) RegisterUser(user *model.User) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}

	user.Password = string(encryptedPassword)
	result := config.DB.Create(user)
	return result.Error
}

func (s *AuthServiceImpl) CheckUserExistence(username string) (bool, error) {
	return common.CheckUserExistence(username)
}

func (s *AuthServiceImpl) RespondWithToken(c *gin.Context, user *model.User) {
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

func (s *AuthServiceImpl) GetUserByUsernameOrEmail(c *gin.Context, usernameOrEmail string) (*model.User, error) {
	user, err := common.GetUserByUsernameOrEmail(usernameOrEmail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":     "Failed to retrieve user",
			"status":      "error",
			"description": "",
		})
		return nil, err
	}

	if user == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":      "error",
			"message":     "Invalid username or email",
			"description": "username or email",
		})
		return nil, fmt.Errorf("invalid username or email")
	}

	return user, nil
}
