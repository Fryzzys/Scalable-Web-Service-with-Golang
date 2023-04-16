package repository

import (
	"project-gorm/models/entity"
)

type BookRepository interface {
	GetBook() ([]entity.Book, error)
	GetBookID(BookID int64) (entity.Book, error)
	CreateBook(book entity.Book) (entity.Book, error)
	UpdateBook(BookID int64, book entity.Book) (entity.Book, error)
	DeleteBook(BookID int64) (string, error)
}