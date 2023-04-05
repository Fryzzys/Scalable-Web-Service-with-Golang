package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Books struct {
	ID     int    `json:"book_id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var (
	db  *sql.DB
	err error
)

func main() {
	db, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=1 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")

	r := gin.Default()

	r.POST("/books", CreateBook)
	r.GET("/books", GetBook)
	r.GET("/books/:ID", GetBook)
	r.PUT("/books/:ID", UpdateBook)
	r.DELETE("/books/:ID", DeleteBook)
	r.Run(":8080")
}

func CreateBook(ctx *gin.Context) {
	var book Books
	query := `
	INSERT INTO books (title, author, describ)
	VALUES ($1, $2, $3)
	RETURNING *
  `

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
	ID, _ := strconv.Atoi(ctx.Param("ID"))
	query := `SELECT * FROM books WHERE id=$1`
	querys := `SELECT * FROM books`

	switch {
	case ID > 0:
		err = db.QueryRow(query, &ID).Scan(&book.ID, &book.Title, &book.Author, &book.Desc)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		results = append(results, book)
		ctx.JSON(http.StatusOK, results)
	case ID == 0:
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
	default:
		panic(err)
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
	err = ctx.ShouldBindJSON(&book)
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
	_, err = db.Exec(query, &ID)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, fmt.Sprint("Data Deleted"))
}
