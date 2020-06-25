package service

import (
	"errors"

	"algogrit.com/go-library/pkg/repository"
)

type BooksServiceImpl struct {
	repository repository.BooksRepository
}

func NewBookService(bookRepo repository.BooksRepository) BooksService {
	return &BooksServiceImpl{bookRepo}
}

func (bs *BooksServiceImpl) Find(Id int) (repository.Book, error) {
	return bs.repository.Select(Id)
}

func (bs *BooksServiceImpl) FindAll() ([]repository.Book, error) {
	books := bs.repository.SelectAll()

	if books == nil {
		return nil, errors.New("No books found!")
	}

	return books, nil
}
