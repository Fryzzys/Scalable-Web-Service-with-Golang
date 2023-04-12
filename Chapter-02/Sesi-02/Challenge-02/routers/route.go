package routers

import (
	"simple-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine{
	r := gin.Default()

	r.POST("/books", controllers.CreateBook)
	r.GET("/books", controllers.GetBook)
	r.GET("/books/:bookID", controllers.GetBook)
	r.PUT("/books/:bookID", controllers.UpdateBook)
	r.DELETE("/books/:bookID", controllers.DeleteBook)

	return r
}