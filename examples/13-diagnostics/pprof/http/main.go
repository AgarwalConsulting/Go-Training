package main

import (
	"fmt"
	"log"
	"net/http"

	_ "net/http/pprof"

	"algogrit.com/pprof/http-demo/blog"
	"github.com/gorilla/mux"
)

var port = "8000"

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/posts", blog.GetPostsHandler).Methods("GET")

	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	fmt.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
