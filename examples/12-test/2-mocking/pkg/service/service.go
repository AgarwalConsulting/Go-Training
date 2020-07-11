package service

import "algogrit.com/go-library/pkg/repository"

type BooksService interface {
	Find(int) (repository.Book, error)
	FindAll() ([]repository.Book, error)
}
