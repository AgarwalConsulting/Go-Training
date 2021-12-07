package repository

import (
	"biblioteca/entities"
	"fmt"
	"sync"
)

func sequence(initValue int) func() int {
	i := initValue

	return func() int {
		i++
		return i
	}
}

type bookInMem struct {
	nextID func() int
	books  map[int]entities.Book
	mut    sync.RWMutex
}

func (repo *bookInMem) FindAll() ([]entities.Book, error) {
	repo.mut.RLock()
	defer repo.mut.RUnlock()

	var books = make([]entities.Book, 0, len(repo.books))

	for _, book := range repo.books {
		books = append(books, book)
	}

	return books, nil
}

func (repo *bookInMem) FindBy(id int) (book entities.Book, err error) {
	repo.mut.RLock()
	defer repo.mut.RUnlock()

	book, ok := repo.books[id]

	if !ok {
		err = fmt.Errorf("Unable to find book by id: %d", id)
	}

	return
}

func (repo *bookInMem) Create(newBook entities.Book) (*entities.Book, error) {
	repo.mut.Lock()
	defer repo.mut.Unlock()

	newBook.ID = repo.nextID()
	repo.books[newBook.ID] = newBook

	return &newBook, nil
}

func (repo *bookInMem) Update(existingBook *entities.Book) error {
	repo.mut.Lock()
	defer repo.mut.Unlock()

	_, ok := repo.books[existingBook.ID]

	if !ok {
		return fmt.Errorf("Unable to locate book:", existingBook.ID)
	}

	repo.books[existingBook.ID] = *existingBook

	return nil
}

func (repo *bookInMem) Destroy(existingBook *entities.Book) error {
	repo.mut.Lock()
	defer repo.mut.Unlock()

	_, ok := repo.books[existingBook.ID]

	if !ok {
		return fmt.Errorf("Unable to locate book:", existingBook.ID)
	}

	delete(repo.books, existingBook.ID)

	return nil
}

func NewInmemRepository() BookRepository {
	nextID := sequence(0)

	bookrepo := bookInMem{nextID: nextID, books: map[int]entities.Book{
		1: {ID: nextID(), Title: "The C Book", Author: "Dennis Ritchie"},
		2: {ID: nextID(), Title: "C++", Author: "Bjarne Stroustrop"},
	}}

	return &bookrepo
}
