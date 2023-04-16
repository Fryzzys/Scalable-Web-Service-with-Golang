package service

import "project-gorm/models/usrRequest"

type BookService interface {
	GetBook() ([]usrRequest.BookResponse, error)
	GetBookID(BookID int64) (usrRequest.BookResponse, error)
	CreateBook(book usrRequest.BookRequest) (usrRequest.BookResponse, error)
	UpdateBook(BookID int64, book usrRequest.BookRequest) (usrRequest.BookResponse, error)
	DeleteBook(BookID int64) (string, error)
}