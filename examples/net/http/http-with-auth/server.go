package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"gopkg.in/go-playground/validator.v10"
)

func employeesIndexHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, employees)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func employeeCreateHandler(w http.ResponseWriter, r *http.Request) {
	var e Employee
	json.NewDecoder(r.Body).Decode(&e)

	e.ID = nextID()

	validate := validator.New()
	errs := validate.Struct(e)

	if errs != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Unable to process: %v", errs)
		return
	}

	employees = append(employees, e)

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(e)
}

func employeeShowHandler(w http.ResponseWriter, r *http.Request) {
	empID := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(empID)

	emp := FindEmployeeBy(id)

	json.NewEncoder(w).Encode(emp)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/healthz", healthzHandler).Methods("GET")
	r.HandleFunc("/employees", employeesIndexHandler).Methods("GET")
	r.HandleFunc("/employees", employeeCreateHandler).Methods("POST")
	r.HandleFunc("/employees/{id}", employeeShowHandler).Methods("GET")

	n := negroni.Classic() // Includes some default middlewares
	n.Use(negroni.HandlerFunc(basicAuthMiddleware))
	n.UseHandler(r)

	http.ListenAndServe(":5000", n)
}
