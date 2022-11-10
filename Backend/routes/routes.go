package routes

import (
	AuthController "Backend/go-api/controller/auth"
	UserController "Backend/go-api/controller/user"
	"Backend/go-api/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/login", AuthController.Login)
	r.POST("/register", AuthController.Register)
	//users
	authorized := r.Group("/users", middleware.ValidationUsers())
	authorized.GET("/get", UserController.GetUser)
	authorized.GET("/profile", UserController.Profile)

	r.Run()
}
