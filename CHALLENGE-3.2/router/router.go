package router

import (
	"CHALLENGE-3.2/controllers"
	"CHALLENGE-3.2/middlewares"
	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.GET("/", controllers.GetAllProduct)

		productRouter.GET("/:productId", middlewares.ProductAuthorization(), controllers.GetBookById)
		productRouter.PUT("/:productId", middlewares.AdminAuthorization(), controllers.UpdateProduct)
		productRouter.DELETE("/:productId", middlewares.AdminAuthorization(), controllers.DeleteProduct)
	}
	return r
}
