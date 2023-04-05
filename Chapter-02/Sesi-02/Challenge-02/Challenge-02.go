package main

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

var BookDatas = []Books{}

func main() {
	r := gin.Default()
	r.POST("/books", CreateBook)
	r.GET("/books", GetBook)
	r.GET("/books/:bookID", GetBook)
	r.PUT("/books/:bookID", UpdateBook)
	r.DELETE("/books/:bookID", DeleteBook)
	r.Run(":8080")
}

func CreateBook(ctx *gin.Context) {
	var newBook Books

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	newBook.ID = len(BookDatas) + 1
	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusCreated, fmt.Sprint("Created"))
}

func GetBook(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("bookID"))
	condition := false
	var bookData Books
	var booksData []Books

	for i, book := range BookDatas {
		if ID == book.ID {
			condition = true
			bookData = BookDatas[i]
			ctx.JSON(http.StatusOK, bookData)
			break
		}
		if ID == 0 {
			condition = true
			booksData = BookDatas
			ctx.JSON(http.StatusOK, booksData)
			break
		}
	}

	if !condition {
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

	for i, book := range BookDatas {
		if ID == book.ID {
			condition = true
			BookDatas[i] = updatedBook
			BookDatas[i].ID = ID
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

	for i, book := range BookDatas {
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
	copy(BookDatas[bookIndex:], BookDatas[bookIndex+1:])
	BookDatas[len(BookDatas)-1] = Books{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK, fmt.Sprint("Deleted"))
}
