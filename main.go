package main

import (
	"Backend/go-api/config"
	"Backend/go-api/routes"

	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	config.Database()
	routes.Router()
}