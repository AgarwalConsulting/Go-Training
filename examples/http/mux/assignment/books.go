package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID          int
	Title       string
	Author      string
	ISBN        string
	Description string
	Price       float64
}

var books = map[int]Book{
	1: Book{ID: 1, Title: "The C Book", Author: "Dennis Ritchie"},
	2: Book{ID: 2, Title: "C++", Author: "Bjarne Stroustrop"},
}

var counter = 3

func booksIndexHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books", booksIndexHandler).Methods("GET")

	http.ListenAndServe(":9000", r)
}
