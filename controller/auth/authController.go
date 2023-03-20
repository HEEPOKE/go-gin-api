package auth

import (
	"Backend/go-api/common"
	"Backend/go-api/model"
	"Backend/go-api/services/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	AuthService auth.AuthService
}

const SecretKey = "secret"

func (a *Auth) Register(c *gin.Context) {
	var json model.User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userExists, _ := a.AuthService.CheckUserExistence(json.Username); userExists {
		c.JSON(http.StatusOK, gin.H{
			"message": "user Exist",
			"status":  "error",
		})
		return
	}

	user := model.User{
		Username: json.Username,
		Password: json.Password,
		Email:    json.Email,
		Tel:      json.Tel,
		Role:     model.Viewer,
	}

	if err := a.AuthService.RegisterUser(&user); err != nil {
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

func (a *Auth) Login(c *gin.Context) {
	var json model.Auth
	err := common.BindJSON(c, &json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := a.AuthService.GetUserByUsernameOrEmail(c, json.UsernameOrEmail)
	if err != nil {
		return
	}

	if !common.ComparePasswords(c, user.Password, json.Password) {
		return
	}

	a.AuthService.RespondWithToken(c, user)
}

func Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout Success",
	})
}
