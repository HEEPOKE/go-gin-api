package main

import (
	"Backend/go-api/config"
	AuthController "Backend/go-api/controller/auth"
	UserController "Backend/go-api/controller/user"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	config.Database()
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/login", AuthController.Login)
	r.POST("/register", AuthController.Register)
	r.GET("/users", UserController.GetUser)
	r.Run()
}
