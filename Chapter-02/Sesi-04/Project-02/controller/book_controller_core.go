package controller

import (
	"project-gorm/models/usrRequest"
	"project-gorm/service"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type BookControllerCore struct {
	BookService service.BookService
}

func NewBookController(bookService service.BookService) BookController {
	return &BookControllerCore {
		BookService: bookService,
	}
}

func (controller *BookControllerCore) GetBook(ctx *gin.Context) {
	result, err := controller.BookService.GetBook()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, usrRequest.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (controller *BookControllerCore) GetBookID(ctx *gin.Context) {
	var BookIDstr = ctx.Param("book_id")

	BookID, err := strconv.Atoi(BookIDstr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, usrRequest.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	result, err := controller.BookService.GetBookID(int64(BookID))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, usrRequest.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, usrRequest.BookResponse{
		BookID:    result.BookID,
		BookName:  result.BookName,
		Author:    result.Author,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	})
}

func (controller *BookControllerCore) CreateBook(ctx *gin.Context) {
	var newBook usrRequest.BookRequest

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, usrRequest.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	result, err := controller.BookService.CreateBook(newBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, usrRequest.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, usrRequest.BookResponse{
		BookID:    result.BookID,
		BookName:  result.BookName,
		Author:    result.Author,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	})
}

func (controller *BookControllerCore) UpdateBook(ctx *gin.Context) {
	var newBook usrRequest.BookRequest
	var BookIDstr = ctx.Param("book_id")

	bookID, err := strconv.Atoi(BookIDstr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, usrRequest.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, usrRequest.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	result, err := controller.BookService.UpdateBook(int64(bookID), newBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, usrRequest.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, usrRequest.BookResponse{
		BookID:    result.BookID,
		BookName:  result.BookName,
		Author:    result.Author,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	})
}

func (controller *BookControllerCore) DeleteBook(ctx *gin.Context) {
	var BookIDstr = ctx.Param("book_id")

	bookID, err := strconv.Atoi(BookIDstr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, usrRequest.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	result, err := controller.BookService.DeleteBook(int64(bookID))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, usrRequest.ErrorResponse{
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}