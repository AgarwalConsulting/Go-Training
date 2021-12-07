package repository

import "biblioteca/entities"

type BookRepository interface {
	FindAll() ([]entities.Book, error)
	FindBy(id int) (entities.Book, error)
	Create(newBook entities.Book) (*entities.Book, error)
	Update(existingBook *entities.Book) error
	Destroy(existingBook *entities.Book) error
}
