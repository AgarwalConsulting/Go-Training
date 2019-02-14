package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var mul int
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			mul = i * j
		}
	}
	fmt.Fprint(w, "Hello HTTP!")
	fmt.Fprint(w, mul)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	// r.HandleFunc("/products", ProductsHandler)
	// r.HandleFunc("/articles", ArticlesHandler)
	// http.Handle("/", r)

	http.ListenAndServe(":8090", r)
}
