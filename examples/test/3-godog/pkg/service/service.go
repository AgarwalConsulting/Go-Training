package service

import "github.com/Chennai-Golang/101-workshop/examples/test/4-mocking/go-library/pkg/repository"

type BooksService interface {
	Find(int) (repository.Book, error)
	FindAll() ([]repository.Book, error)
}
