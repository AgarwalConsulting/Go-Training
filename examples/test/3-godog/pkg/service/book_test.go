package service

import (
	"testing"

	"github.com/Chennai-Golang/101-workshop/examples/test/4-mocking/go-library/pkg/repository"
	"github.com/stretchr/testify/assert"
)

type mockBookRepository struct {
	_select    func(int) (repository.Book, error)
	_selectAll func() []repository.Book
}

func (mockBR mockBookRepository) Select(Id int) (repository.Book, error) {
	if mockBR._select != nil {
		return mockBR._select(Id)
	}
	return repository.Book{}, nil
}

func (mockBR mockBookRepository) SelectAll() []repository.Book {
	if mockBR._selectAll != nil {
		return mockBR._selectAll()
	}

	return nil
}

func TestBookServiceSelectAll(t *testing.T) {
	expectedBook := repository.Book{Title: "Alchemist"}

	mockBR := mockBookRepository{
		_selectAll: func() []repository.Book {
			return []repository.Book{
				expectedBook,
			}
		},
	}

	bookServiceUT := NewBookService(mockBR)

	books, err := bookServiceUT.FindAll()

	assert.Nil(t, err)

	assert.Contains(t, books, expectedBook)
}
