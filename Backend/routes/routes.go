package routes

import (
	AuthController "Backend/go-api/controller/auth"
	ProductController "Backend/go-api/controller/product"
	UserController "Backend/go-api/controller/user"
	"Backend/go-api/middleware"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("ENDPOINT_URL")},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == os.Getenv("ENDPOINT_URL")
		},
		MaxAge: 12 * time.Hour,
	}))
	auth := r.Group("/api/auth")
	{
		auth.GET("/login", AuthController.Login)
		auth.POST("/register", AuthController.Register)
		auth.GET("/logout", AuthController.Logout)
	}
	//users
	authorized := r.Group("/api/users", middleware.ValidationUsers())
	{
		authorized.GET("/get", UserController.GetUser)
		authorized.GET("/get/:id", UserController.GetUserById)
		authorized.GET("/profile", UserController.Profile)
	}
	// product
	product := r.Group("/api/product")
	{
		product.GET("/get", ProductController.ReadProduct)
		product.POST("/create", ProductController.Create)
		product.PUT("/edit/:id", ProductController.Edit)
		product.DELETE("/delete/:id", ProductController.Delete)
	}

	r.Run()
}
