package repository

import (
	"errors"
	"fmt"
	"sync"
)

type Book struct {
	ID          int
	Title       string
	Author      string
	ISBN        string
	Description string
	Price       float64
}

type BooksInMapRepository struct {
	books map[int]Book
	lock  sync.RWMutex
}

func NewBooksRepository() BooksRepository {
	bookMap := map[int]Book{
		1: Book{ID: 1, Title: "The C Book", Author: "Dennis Ritchie"},
		2: Book{ID: 2, Title: "C++", Author: "Bjarne Stroustrop"},
	}

	return &BooksInMapRepository{books: bookMap}
}

func (br *BooksInMapRepository) SelectAll() []Book {

	books := []Book{}

	for _, v := range br.books {
		books = append(books, v)
	}

	return books
}

func (br *BooksInMapRepository) Select(Id int) (Book, error) {
	book, ok := br.books[Id]

	if ok {
		return book, nil
	} else {
		return book, errors.New(fmt.Sprintf("Book with id: %d not found!", Id))
	}
}
