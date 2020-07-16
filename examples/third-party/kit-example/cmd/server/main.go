package main

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"

	healthService "algogrit.com/emp-server/health/service"

	employeeEndpoints "algogrit.com/emp-server/employee/endpoints"
	employeeRepository "algogrit.com/emp-server/employee/repository"
	employeeService "algogrit.com/emp-server/employee/service"
	employeeTransport "algogrit.com/emp-server/employee/transport"
)

func main() {
	r := mux.NewRouter()

	hService := healthService.New()

	r.HandleFunc("/healthz", hService.HealthInfo).Methods("GET")

	eRepo := employeeRepository.NewInMem()
	eService := employeeService.NewService(eRepo)
	eEndpoints := employeeEndpoints.New(eService)

	eHandler := employeeTransport.NewHTTPServer(eEndpoints)

	go http.ListenAndServe(":8000", eHandler)

	// r.Handle("/employees", auth.BasicAuthMiddleware(http.HandlerFunc(eService.Index))).Methods("GET")
	// r.Handle("/employees", auth.BasicAuthMiddleware(http.HandlerFunc(eService.Create))).Methods("POST")
	// r.Handle("/employees/{id}", auth.BasicAuthMiddleware(http.HandlerFunc(eService.Show))).Methods("GET")

	n := negroni.Classic() // Includes some default middlewares
	// n.Use(negroni.HandlerFunc(auth.BasicAuthMiddleware))
	n.UseHandler(r)

	err := http.ListenAndServe(":9000", n)
	log.Error(err)
}
