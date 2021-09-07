package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Employee struct {
	ID         int    `json:"-"`
	Name       string `json:"name"`
	Department string `json:"speciality"`
	ProjectID  int    `json:"project"`
}

// func (e Employee) MarshalJSON() ([]byte, error) {
// 	jsonString := fmt.Sprintf(`{"name": "%s", "speciality": "%s", "project": %d}`, e.Name, e.Department, e.ProjectID)

// 	// fmt.Println(jsonString)

// 	return []byte(jsonString), nil
// }

var employees = []Employee{
	{1, "Gaurav Agarwal", "LnD", 1001},
	{2, "Mohamed", "Cloud", 1010},
	{3, "Suganthi", "DBA", 1002},
	{4, "Vishnu", "Serverless", 1011},
}

// func EmployeesHandler(w http.ResponseWriter, req *http.Request) {
// 	if req.Method == "POST" {
// 		EmployeeCreateHandler(w, req)
// 	} else if req.Method == "GET" {
// 		EmployeesIndexHandler(w, req)
// 	} else {
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		return
// 	}
// }

func EmployeesIndexHandler(w http.ResponseWriter, req *http.Request) {
	// w.Header().Add("Content-Type", "application/json")
	// w.Header().Add("Content-Type", "charset=utf-8")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// fmt.Fprintln(w, employees)
	encoder := json.NewEncoder(w)
	encoder.Encode(employees)
}

func EmployeeCreateHandler(w http.ResponseWriter, req *http.Request) {
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
	// r := http.NewServeMux()
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!"

		// w.Write([]byte(msg))
		fmt.Fprintln(w, msg)
	})

	r.HandleFunc("/employees", EmployeesIndexHandler).Methods("GET")
	r.HandleFunc("/employees", EmployeeCreateHandler).Methods("POST")
	// r.HandleFunc("/employees", EmployeesHandler)

	http.ListenAndServe(":8000", r)
}
