package controllers

import (
	"fmt"
	"net/http"
	"simple-rest-api-with-db/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Books struct {
	ID     int    `json:"book_id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

func CreateBook(ctx *gin.Context) {
	var book Books
	query := `
	INSERT INTO books (title, author, describ)
	VALUES ($1, $2, $3)
	RETURNING *`

	db := database.DbConnection()

	err := ctx.ShouldBindJSON(&book)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = db.Exec(query, book.Title, book.Author, book.Desc)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, fmt.Sprint("Data Created"))
}

func GetBook(ctx *gin.Context) {
	var results = []Books{}
	var book Books
	idString := ctx.Param("ID")
	ID, _ := strconv.Atoi(idString)
	query := `SELECT * FROM books WHERE id=$1`
	querys := `SELECT * FROM books`

	db := database.DbConnection()

	if idString != "" {
		err := db.QueryRow(query, &ID).Scan(
			&book.ID, &book.Title, &book.Author, &book.Desc)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Book with id: %d not found", ID),
			})
			return
		}
		results = append(results, book)
		ctx.JSON(http.StatusOK, results)
	}
	if idString == "" {
		row, err := db.Query(querys)
		if err != nil {
			panic(err)
		}
		for row.Next() {
			err = row.Scan(&book.ID, &book.Title, &book.Author, &book.Desc)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			results = append(results, book)
		}
		ctx.JSON(http.StatusOK, results)
	}
}

func UpdateBook(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("ID"))
	var book Books
	query := `
	UPDATE books
	SET title = $2, author = $3, describ = $4
	WHERE id=$1;
	`
	db := database.DbConnection()

	err := ctx.ShouldBindJSON(&book)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	_, err = db.Exec(query, &ID, &book.Title, &book.Author, &book.Desc)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotModified, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, fmt.Sprint("Data Updated"))
}

func DeleteBook(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("ID"))
	query := `
	DELETE FROM books
	WHERE id=$1;
	`
	db := database.DbConnection()

	_, err := db.Exec(query, &ID)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, fmt.Sprint("Data Deleted"))
}
