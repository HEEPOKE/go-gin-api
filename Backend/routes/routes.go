package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func router() {
	r := gin.Default()
	r.Use(cors.Default())

}
