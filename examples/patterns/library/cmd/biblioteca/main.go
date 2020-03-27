package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"algogrit.com/library/pkg/entity"
)

type Library struct {
	Books   map[int]entity.Book
	counter func() int
}

func sequence(initValue int) func() int {
	i := initValue

	return func() int {
		i++
		return i
	}
}

// GET /books/{id}
func (lib Library) bookShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID, _ := strconv.Atoi(vars["id"])

	log.Info("Getting bookID: ", bookID)

	book := lib.Books[bookID]

	json.NewEncoder(w).Encode(book)
}

// GET /books
func (lib Library) booksIndexHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(lib.Books)
}

func NewLibrary() Library {
	var books = map[int]entity.Book{
		1: entity.Book{ID: 1, Title: "The C Book", Author: "Dennis Ritchie"},
		2: entity.Book{ID: 2, Title: "C++", Author: "Bjarne Stroustrop"},
	}

	var counter = sequence(2)

	return Library{books, counter}
}

func main() {
	port := "9000"
	r := mux.NewRouter()

	library := NewLibrary()

	r.HandleFunc("/books", library.booksIndexHandler).Methods("GET")
	r.HandleFunc("/books/{id}", library.bookShowHandler).Methods("GET")

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(r)

	log.Info("Server is running on port: ", port)

	http.ListenAndServe(":"+port, n)
}
