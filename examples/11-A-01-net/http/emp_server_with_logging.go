package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Employee struct {
	ID         int    `json:"-"`
	Name       string `json:"name"`
	Department string `json:"speciality"`
	ProjectID  int    `json:"project"`
}

var employees = []Employee{
	{1, "Gaurav Agarwal", "LnD", 1001},
	{2, "Mohamed", "Cloud", 1010},
	{3, "Suganthi", "DBA", 1002},
	{4, "Vishnu", "Serverless", 1011},
}

func EmployeesIndexHandler(w http.ResponseWriter, req *http.Request) {
	// log.Printf("%s %s: EmployeesIndexHandler\n", req.Method, req.URL)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	encoder := json.NewEncoder(w)
	encoder.Encode(employees)
}

func EmployeeCreateHandler(w http.ResponseWriter, req *http.Request) {
	// log.Printf("%s %s: EmployeeCreateHandler\n", req.Method, req.URL)

	var newEmployee Employee

	err := json.NewDecoder(req.Body).Decode(&newEmployee)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	newEmployee.ID = len(employees) + 1

	employees = append(employees, newEmployee)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(newEmployee)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!"

		fmt.Fprintln(w, msg)
	})

	r.HandleFunc("/employees", EmployeesIndexHandler).Methods("GET")
	r.HandleFunc("/employees", EmployeeCreateHandler).Methods("POST")

	http.ListenAndServe(":8000", handlers.CombinedLoggingHandler(os.Stdout, r))
}
