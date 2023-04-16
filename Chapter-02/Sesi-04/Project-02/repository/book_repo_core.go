package repository

import (
	"errors"
	"fmt"
	"project-gorm/models/entity"
	"gorm.io/gorm"
)

type BookRepo struct {
	DB *gorm.DB
}

func NewBookRepo(db *gorm.DB) BookRepository {
	return &BookRepo{
		DB: db,
	}
}

func (repository *BookRepo) GetBook() ([]entity.Book, error) {
	results := []entity.Book{}

	err := repository.DB.Find(&results).Error
	if err != nil {
		return []entity.Book{}, err
	}

	return results, nil
}

func (repository *BookRepo) GetBookID(BookID int64) (entity.Book, error) {
	results := entity.Book{}

	err := repository.DB.First(&results, "book_id = ?", BookID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Book{}, errors.New(fmt.Sprintf("Book with id %d not found", BookID))
		}

		return entity.Book{}, err
	}

	return results, nil
}

func (repository *BookRepo) CreateBook(book entity.Book) (entity.Book, error) {
	newBook := entity.Book{
		BookName: book.BookName,
		Author:   book.Author,
	}

	err := repository.DB.Create(&newBook).Error
	if err != nil {
		return entity.Book{}, err
	}

	return newBook, nil
}

func (repository *BookRepo) UpdateBook(BookID int64, book entity.Book) (entity.Book, error) {
	bookUpdated := entity.Book{}

	err := repository.DB.First(&bookUpdated, "book_id = ?", BookID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Book{}, errors.New(fmt.Sprintf("Book with id %d not found", BookID))
		}

		return entity.Book{}, err
	}

	err = repository.DB.Model(&bookUpdated).Updates(entity.Book{BookName: book.BookName, Author: book.Author}).Error
	if err != nil {
		return entity.Book{}, err
	}

	return bookUpdated, nil
}

func (repository *BookRepo) DeleteBook(BookID int64) (string, error) {
	bookDeleted := entity.Book{}

	err := repository.DB.First(&bookDeleted, "book_id = ?", BookID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New(fmt.Sprintf("Book with id %d not found", BookID))
		}

		return "", err
	}

	err = repository.DB.Delete(&bookDeleted, BookID).Error
	fmt.Println("err", err)
	if err != nil {
		return "", err
	}

	return "Book deleted successfully", nil
}