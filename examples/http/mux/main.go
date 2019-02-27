package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type Person struct {
	Name   string
	Age    int
	Income int `json:"-"`
}

func helloWorldHandler(w http.ResponseWriter, req *http.Request) {
	name := mux.Vars(req)["name"]
	person := Person{Name: name}
	json.NewEncoder(w).Encode(&person)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{name}", helloWorldHandler).Methods("GET")

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(r)

	http.ListenAndServe(":9000", n)
}
