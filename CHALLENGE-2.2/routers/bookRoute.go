package routers

import (
	"CHALLENGE-2.2/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.GET("/books/", controllers.GetAllBook)
	router.POST("/books", controllers.CreateBook)
	router.PUT("/books/:bookID", controllers.UpdateBook)
	router.GET("/books/:bookID", controllers.GetBook)
	router.DELETE("/books/:bookID", controllers.DeleteBook)
	return router
}
