package main

import (
	"Backend/go-api/config"
	"Backend/go-api/controller"
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
	r.POST("/login", controller.Login)
	r.POST("/register", controller.Register)
	r.Run()
}
