package routes

import (
	AuthController "Backend/go-api/controller/auth"
	UserController "Backend/go-api/controller/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/login", AuthController.Login)
	r.POST("/register", AuthController.Register)
	r.GET("/users", UserController.GetUser)
	r.Run()
}
