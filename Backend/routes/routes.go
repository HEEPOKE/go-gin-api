package routes

import (
	AuthController "Backend/go-api/controller/auth"
	ProductController "Backend/go-api/controller/product"
	UserController "Backend/go-api/controller/user"
	"Backend/go-api/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/api/login", AuthController.Login)
	r.POST("/api/register", AuthController.Register)
	//users
	authorized := r.Group("/api/users", middleware.ValidationUsers())
	{
		authorized.GET("/get", UserController.GetUser)
		authorized.GET("/get/:id", UserController.LockUser)
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
