package main

import (
	"log"
	"net/http"

	_ "net/http/pprof"

	"github.com/Chennai-Golang/101-workshop/examples/diagnostics/pprof/http/blog"
	"github.com/gorilla/mux"
)

var port = "8000"

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/posts", blog.GetPostsHandler).Methods("GET")

	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	log.Fatal(http.ListenAndServe(":"+port, r))
}
