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
	r.POST("/login", AuthController.Login)
	r.POST("/register", AuthController.Register)
	//users
	authorized := r.Group("/users", middleware.ValidationUsers())
	authorized.GET("/get", UserController.GetUser)
	authorized.GET("/get/:id", UserController.LockUser)
	authorized.GET("/profile", UserController.Profile)
	// product
	r.POST("/product/create", ProductController.Create)
	r.POST("/product/edit/:id", ProductController.Edit)
	r.Run()
}
