package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func helloWorldHandler(w http.ResponseWriter, req *http.Request) {
	msg := []byte("Hello, World!")
	w.Write(msg)
}

func goodbyeWorldHandler(w http.ResponseWriter, req *http.Request) {
	msg := []byte("Bye, World!")
	w.Write(msg)
}

func main() {
	r := mux.NewRouter()

	n := negroni.Classic()

	r.HandleFunc("/", helloWorldHandler).Methods("GET")
	r.HandleFunc("/bye", helloWorldHandler).Methods("POST")
	// http.HandleFunc("/", helloWorldHandler)
	// http.HandleFunc("/bye", goodbyeWorldHandler)

	n.UseHandler(r)

	http.ListenAndServe(":9000", n)
}
