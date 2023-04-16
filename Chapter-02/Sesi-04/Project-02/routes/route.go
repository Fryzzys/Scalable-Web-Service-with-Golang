package routes

import (
	"project-gorm/controller"
	"project-gorm/repository"
	"project-gorm/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupBookRoute(router *gin.Engine, db *gorm.DB) {
	bookRepository := repository.NewBookRepo(db)
	bookService := service.NewBookService(bookRepository)
	bookController := controller.NewBookController(bookService)

	router.POST("/books", bookController.CreateBook)
	router.GET("/books", bookController.GetBook)
	router.GET("/books/:book_id", bookController.GetBookID)
	router.PUT("/books/:book_id", bookController.UpdateBook)
	router.DELETE("/books/:book_id", bookController.DeleteBook)
}
