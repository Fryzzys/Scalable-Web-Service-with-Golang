package routers

import (
	"simple-rest-api-with-db/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()

	r.POST("/books", controllers.CreateBook)
	r.GET("/books", controllers.GetBook)
	r.GET("/books/:ID", controllers.GetBook)
	r.PUT("/books/:ID", controllers.UpdateBook)
	r.DELETE("/books/:ID", controllers.DeleteBook)
	
	return r
}