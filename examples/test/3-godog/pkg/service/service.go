package service

import "github.com/AgarwalConsulting/Go-Training/examples/test/4-mocking/go-library/pkg/repository"

type BooksService interface {
	Find(int) (repository.Book, error)
	FindAll() ([]repository.Book, error)
}
