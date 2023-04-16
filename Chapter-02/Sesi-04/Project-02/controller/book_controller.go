package controller

import "github.com/gin-gonic/gin"

type BookController interface {
	GetBookID(ctx *gin.Context)
	GetBook(ctx *gin.Context)
	CreateBook(ctx *gin.Context)
	UpdateBook(ctx *gin.Context)
	DeleteBook(ctx *gin.Context)
}