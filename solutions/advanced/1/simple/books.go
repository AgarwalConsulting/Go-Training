package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
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
	fmt.Println("-Executing booksIndex Handler...")
	defer fmt.Println("-Completed booksIndex Handler...")
	mut.RLock()
	defer mut.RUnlock()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func bookIDMiddleware(next http.HandlerFunc) http.HandlerFunc {
	bookHandler := func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("--Request intercepted: bookIDMiddleware...")
		defer fmt.Println("--Request completed: bookIDMiddleware...")

		mut.RLock()

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

		ctx := req.Context()

		ctxWithBook := context.WithValue(ctx, "currentBook", book)
		ctxWithBook = context.WithValue(ctxWithBook, "HasBook", true)

		childReq := req.WithContext(ctxWithBook)
		mut.RUnlock()

		next(w, childReq)
	}

	return http.HandlerFunc(bookHandler)
}

// func getCurrentBook(w http.ResponseWriter, req *http.Request) *Book {
// 	vars := mux.Vars(req)
// 	bookID, err := strconv.Atoi(vars["id"])

// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprintln(w, err)
// 		return nil
// 	}

// 	log.Println("Getting bookID: ", bookID)
// 	book, ok := books[bookID]

// 	if !ok {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprintln(w, "unable to locate book with id:", bookID)
// 		return nil
// 	}

// 	return &book
// }

// GET /books/{id}
func bookShowHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("-Executing bookShow Handler...")
	defer fmt.Println("-Completed bookShow Handler...")
	mut.RLock()
	defer mut.RUnlock()

	book := req.Context().Value("currentBook")

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// POST /books - TODO: Check if a book with the same ISBNCode already exists!
func bookCreateHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("-Executing bookCreate Handler...")
	defer fmt.Println("-Completed bookCreate Handler...")
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
	fmt.Println("-Executing bookCreate Handler...")
	defer fmt.Println("-Completed bookCreate Handler...")
	mut.Lock()
	defer mut.Unlock()

	existingBook := req.Context().Value("currentBook").(Book)

	books[existingBook.ID] = existingBook
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingBook)
}

// DELETE /books/{id}
func bookDeleteHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("-Executing bookDelete Handler...")
	defer fmt.Println("-Completed bookDelete Handler...")
	mut.Lock()
	defer mut.Unlock()

	existingBook := req.Context().Value("currentBook").(Book)

	delete(books, existingBook.ID)

	w.WriteHeader(http.StatusNoContent)
	// w.Header().Add("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(existingBook)
}

// Middleware
func loggingHandler(h http.Handler) http.Handler {
	handlerFn := func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("--Request intercepted: LoggingMiddleware...")
		defer fmt.Println("--Request completed: LoggingMiddleware...")
		beginTime := time.Now()

		h.ServeHTTP(w, req)

		fmt.Println("Request:", req.Method, req.URL, "completed in", time.Since(beginTime))
	}

	return http.HandlerFunc(handlerFn)
}

func main() {
	// var port = flag.Int("port", 8000, "port number to start the server on")
	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8000"
	}

	flag.Parse()

	// port := "9000"
	// r := http.NewServeMux()
	r := mux.NewRouter()

	r.HandleFunc("/books", booksIndexHandler).Methods("GET")
	r.HandleFunc("/books/{id}", bookIDMiddleware(bookShowHandler)).Methods("GET")
	r.HandleFunc("/books", bookCreateHandler).Methods("POST")
	r.HandleFunc("/books/{id}", bookIDMiddleware(bookUpdateHandler)).Methods("PUT", "PATCH")
	r.HandleFunc("/books/{id}", bookIDMiddleware(bookDeleteHandler)).Methods("DELETE")

	// log.Println("Server is running on port: ", *port)
	log.Println("Server is running on port: ", port)

	// handler := cors.Default().Handler(loggingHandler(r))

	// http.ListenAndServe(":"+port, handler)

	// http.ListenAndServe(":"+strconv.Itoa(*port), loggingHandler(r))
	http.ListenAndServe(":"+port, loggingHandler(r))

	// http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, r))
}
