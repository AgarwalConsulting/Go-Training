package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type Person struct {
	Id     int
	Name   string
	Age    int
	Income int `json:"-"`
}

var people = make(map[int]Person)
var idCounter = 0

// GET /people
func getPeopleHandler(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(&people)
}

// GET /people/{id}
func getPersonHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	person := people[int(id)]

	json.NewEncoder(w).Encode(&person)
}

// POST /people
func createPersonHandler(w http.ResponseWriter, req *http.Request) {
	var person Person

	decoder := json.NewDecoder(req.Body)

	if err := decoder.Decode(&person); err != nil {
		fmt.Println("Error in decoding!")
		return
	}

	idCounter = idCounter + 1
	person.Id = idCounter
	people[idCounter] = person

	json.NewEncoder(w).Encode(&person)
}

func main() {
	port := os.Getenv("PORT")
	r := mux.NewRouter()
	r.HandleFunc("/people", getPeopleHandler).Methods("GET")      // Read
	r.HandleFunc("/people/{id}", getPersonHandler).Methods("GET") // Read
	r.HandleFunc("/people", createPersonHandler).Methods("POST")  // Create

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(r)

	http.ListenAndServe(":"+port, n)
}
