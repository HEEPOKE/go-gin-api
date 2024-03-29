package routes

import (
	"Backend/go-api/controller/auth"
	ProductController "Backend/go-api/controller/product"
	UserController "Backend/go-api/controller/user"
	"Backend/go-api/middleware"
	authServices "Backend/go-api/services/auth"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}

	r.Use(cors.New(corsConfig))

	authService := &authServices.AuthServiceImpl{}

	authController := &auth.Auth{AuthService: authService}

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", authController.Login)
		auth.POST("/register", authController.Register)
		auth.GET("/logout", authController.Logout)
	}

	authorized := r.Group("/api/users", middleware.ValidationUsers())
	{
		authorized.GET("/get/:id", UserController.GetUserById)
		authorized.GET("/profile", UserController.Profile)
	}

	product := r.Group("/api/product")
	{
		product.GET("/read", ProductController.ReadProduct)
		product.POST("/create", ProductController.Create)
		product.PUT("/edit/:id", ProductController.Edit)
		product.DELETE("/delete/:id", ProductController.Delete)
	}
	r.GET("/api/read/users", UserController.GetUser)

	host := os.Getenv("LOCALHOST")
	r.Run(host)
}
