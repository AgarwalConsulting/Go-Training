package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
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
	ID          int     `json:"-"`
	Title       string  `json:"title"`
	Author      string  `json:"writer"`
	ISBN        string  `json:"ISBNCode"`
	Description string  `json:"synopsis"`
	Price       float64 `json:"price"`
}

var nextID = sequence(0)

var books = map[int]Book{
	1: {ID: nextID(), Title: "The C Book", Author: "Dennis Ritchie"},
	2: {ID: nextID(), Title: "C++", Author: "Bjarne Stroustrop"},
}
var mut sync.RWMutex

// GET /books
func booksIndexHandler(w http.ResponseWriter, req *http.Request) {
	mut.RLock()
	defer mut.RUnlock()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GET /books/{id}
func bookShowHandler(w http.ResponseWriter, req *http.Request) {
	mut.RLock()
	defer mut.RUnlock()
	vars := mux.Vars(req)
	bookID, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	log.Println("Getting bookID: ", bookID)
	book, ok := books[bookID]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "unable to locate book with id:", bookID)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// POST /books - TODO: Check if a book with the same ISBNCode already exists!
func bookCreateHandler(w http.ResponseWriter, req *http.Request) {
	mut.Lock()
	defer mut.Unlock()
	var newBook Book
	err := json.NewDecoder(req.Body).Decode(&newBook)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintln(w, err)                 // Unable to read the json req body
		return
	}

	ID := nextID()

	newBook.ID = ID
	books[ID] = newBook

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newBook)
}

// PUT /books/{id}
func bookUpdateHandler(w http.ResponseWriter, req *http.Request) {
	mut.Lock()
	defer mut.Unlock()
	vars := mux.Vars(req)
	bookID, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	existingBook, ok := books[bookID]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "unable to locate book with id:", bookID)
		return
	}

	err = json.NewDecoder(req.Body).Decode(&existingBook)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintln(w, err)
		return
	}

	books[bookID] = existingBook
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingBook)
}

// DELETE /books/{id}
func bookDeleteHandler(w http.ResponseWriter, req *http.Request) {
	mut.Lock()
	defer mut.Unlock()
	vars := mux.Vars(req)
	bookID, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	existingBook, ok := books[bookID]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "unable to locate book with id:", bookID)
		return
	}

	delete(books, bookID)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingBook)
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
	r.HandleFunc("/books", bookCreateHandler).Methods("POST")
	r.HandleFunc("/books/{id}", bookUpdateHandler).Methods("PUT", "PATCH")
	r.HandleFunc("/books/{id}", bookDeleteHandler).Methods("DELETE")

	log.Println("Server is running on port: ", port)

	// handler := cors.Default().Handler(loggingHandler(r))

	// http.ListenAndServe(":"+port, handler)
	http.ListenAndServe(":"+port, loggingHandler(r))
	// http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, r))
}
