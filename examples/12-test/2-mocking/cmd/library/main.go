package main

import (
	"fmt"

	"algogrit.com/go-library/pkg/repository"
	"algogrit.com/go-library/pkg/service"
)

var booksService service.BooksService

func init() {
	bookRepo := repository.NewBooksRepository()

	booksService = service.NewBookService(bookRepo)
}

func main() {
	fmt.Println(booksService.FindAll())
}
