package main

import (
	"fmt"

	"github.com/Chennai-Golang/101-workshop/examples/test/4-mocking/go-library/pkg/repository"
	"github.com/Chennai-Golang/101-workshop/examples/test/4-mocking/go-library/pkg/service"
)

var booksService service.BooksService

func init() {
	bookRepo := repository.NewBooksRepository()

	booksService = service.NewBookService(bookRepo)
}

func main() {
	fmt.Println(booksService.FindAll())
}
