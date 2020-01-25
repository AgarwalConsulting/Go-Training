package repository

type BooksRepository interface {
	Select(int) (Book, error)
	SelectAll() []Book
}
