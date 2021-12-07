package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"log"

	"github.com/gorilla/mux"
)

func sequence(initValue int) func() int {
	i := initValue

	return func() int {
		i++
		return i
	}
}

// Book represents a single book
type Book struct {
	ID          int
	Title       string
	Author      string
	ISBN        string
	Description string
	Price       float64
}

var nextID = sequence(0)

var books = map[int]Book{
	1: Book{ID: nextID(), Title: "The C Book", Author: "Dennis Ritchie"},
	2: Book{ID: nextID(), Title: "C++", Author: "Bjarne Stroustrop"},
}

// GET /books/{id}
func bookShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID, _ := strconv.Atoi(vars["id"])
	log.Println("Getting bookID: ", bookID)
	book := books[bookID]

	json.NewEncoder(w).Encode(book)
}

// GET /books
func booksIndexHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

// Middleware
func loggingHandler(h http.Handler) http.Handler {
	handlerFn := func(w http.ResponseWriter, req *http.Request) {
		beginTime := time.Now()

		h.ServeHTTP(w, req)

		fmt.Println("Request:", req.Method, req.URL, "completed in", time.Since(beginTime))
	}

	return http.HandlerFunc(handlerFn)
}

func main() {
	port := "9000"
	// r := http.NewServeMux()
	r := mux.NewRouter()

	r.HandleFunc("/books", booksIndexHandler).Methods("GET")
	r.HandleFunc("/books/{id}", bookShowHandler).Methods("GET")

	log.Println("Server is running on port: ", port)

	// handler := cors.Default().Handler(loggingHandler(r))

	// http.ListenAndServe(":"+port, handler)
	http.ListenAndServe(":"+port, loggingHandler(r))
	// http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, r))
}
