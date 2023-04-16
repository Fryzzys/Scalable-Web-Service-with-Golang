package service

import (
	"project-gorm/models/entity"
	"project-gorm/models/usrRequest"
	"project-gorm/repository"
)

type BookServiceCore struct {
	BookRepository repository.BookRepository
}

func NewBookService(bookRepository repository.BookRepository) BookService {
	return &BookServiceCore {
		BookRepository: bookRepository,
	}
}

func (service *BookServiceCore) GetBook() ([]usrRequest.BookResponse, error) {
	booksResponse := []usrRequest.BookResponse{}

	result, err := service.BookRepository.GetBook()
	if err != nil {
		return []usrRequest.BookResponse{}, err
	}

	for _, val := range result {
		booksResponse = append(booksResponse, usrRequest.BookResponse(val))
	}

	return booksResponse, nil
}

func (service *BookServiceCore) GetBookID(bookID int64) (usrRequest.BookResponse, error) {
	result, err := service.BookRepository.GetBookID(bookID)
	if err != nil {
		return usrRequest.BookResponse{}, err
	}

	return usrRequest.BookResponse{
		BookID:    result.BookID,
		BookName:  result.BookName,
		Author:    result.Author,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (service *BookServiceCore) CreateBook(
	book usrRequest.BookRequest) (usrRequest.BookResponse, error) {
	newBook := entity.Book{
		BookName: book.BookName,
		Author:   book.Author,
	}

	result, err := service.BookRepository.CreateBook(newBook)
	if err != nil {
		return usrRequest.BookResponse{}, err
	}

	return usrRequest.BookResponse{
		BookID:    result.BookID,
		BookName:  result.BookName,
		Author:    result.Author,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (service *BookServiceCore) UpdateBook(
	bookID int64, book usrRequest.BookRequest) (usrRequest.BookResponse, error) {
	newBook := entity.Book{
		BookName: book.BookName,
		Author:   book.Author,
	}

	result, err := service.BookRepository.UpdateBook(bookID, newBook)
	if err != nil {
		return usrRequest.BookResponse{}, err
	}

	return usrRequest.BookResponse{
		BookID:    result.BookID,
		BookName:  result.BookName,
		Author:    result.Author,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (service *BookServiceCore) DeleteBook(bookID int64) (string, error) {
	result, err := service.BookRepository.DeleteBook(bookID)
	if err != nil {
		return "", err
	}

	return result, nil
}