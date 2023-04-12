package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Books struct {
	ID     int    `json:"book_id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var BookTemp = []Books{}

func CreateBook(ctx *gin.Context) {
	var newBook Books

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	newBook.ID = len(BookTemp) + 1
	BookTemp = append(BookTemp, newBook)

	ctx.JSON(http.StatusCreated, fmt.Sprint("Created"))
}

func GetBook(ctx *gin.Context) {
	idString := ctx.Param("bookID")
	ID, _ := strconv.Atoi(idString)
	condition := false
	var bookData Books
	var booksData []Books

	for i, book := range BookTemp {
		if ID == book.ID {
			condition = true
			bookData = BookTemp[i]
			ctx.JSON(http.StatusOK, bookData)
			break
		}
		if idString == "" {
			condition = true
			booksData = BookTemp
			ctx.JSON(http.StatusOK, booksData)
			break
		}
	}

	if !condition || idString == "0" {
		ctx.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprint("Book not found"))
		return
	}
}

func UpdateBook(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("bookID"))
	condition := false
	var updatedBook Books

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range BookTemp {
		if ID == book.ID {
			condition = true
			BookTemp[i] = updatedBook
			BookTemp[i].ID = ID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("Book with id %v not found", ID))
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprint("Updated"))
}

func DeleteBook(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("bookID"))
	condition := true
	var bookIndex int

	for i, book := range BookTemp {
		if ID == book.ID {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, fmt.Sprintf("Book with id %v not found", ID))
		return
	}
	copy(BookTemp[bookIndex:], BookTemp[bookIndex+1:])
	BookTemp[len(BookTemp)-1] = Books{}
	BookTemp = BookTemp[:len(BookTemp)-1]

	ctx.JSON(http.StatusOK, fmt.Sprint("Deleted"))
}
