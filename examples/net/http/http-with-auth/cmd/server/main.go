package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"algogrit.com/emp-server/auth"
	employeeRepository "algogrit.com/emp-server/employee/repository"
	employeeService "algogrit.com/emp-server/employee/service"
	healthService "algogrit.com/emp-server/health/service"
)

func main() {
	r := mux.NewRouter()

	hService := healthService.New()

	r.HandleFunc("/healthz", hService.HealthInfo).Methods("GET")

	eRepo := employeeRepository.NewInMem()
	eService := employeeService.NewJSON(eRepo)

	r.HandleFunc("/employees", eService.Index).Methods("GET")
	r.HandleFunc("/employees", eService.Create).Methods("POST")
	r.HandleFunc("/employees/{id}", eService.Show).Methods("GET")

	n := negroni.Classic() // Includes some default middlewares
	n.Use(negroni.HandlerFunc(auth.BasicAuthMiddleware))
	n.UseHandler(r)

	http.ListenAndServe(":5000", n)
}
