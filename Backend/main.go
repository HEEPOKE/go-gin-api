package main

import (
	"Backend/go-api/config"
	"Backend/go-api/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Database()
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/register", controller.Register())
	r.Run()
}
