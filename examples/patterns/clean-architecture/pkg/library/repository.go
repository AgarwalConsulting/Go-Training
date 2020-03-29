package library

import "algogrit.com/library/pkg/entity"

//Reader interface
type Reader interface {
	Find(id entity.ID) (*entity.Book, error)
	// Search(query string) ([]*entity.Book, error)
	FindAll() ([]*entity.Book, error)
}

//Writer bookmark writer
type Writer interface {
	Store(b *entity.Book) (entity.ID, error)
	Delete(id entity.ID) error
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase use case interface
type UseCase interface {
	Reader
	Writer
}
