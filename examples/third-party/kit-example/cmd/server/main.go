package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	healthService "algogrit.com/emp-server/health/service"

	employeeEndpoints "algogrit.com/emp-server/employee/endpoints"
	employeeRepository "algogrit.com/emp-server/employee/repository"
	employeeService "algogrit.com/emp-server/employee/service"
	employeeTransport "algogrit.com/emp-server/employee/transport"
)

func main() {
	// Health Service on port :9000
	r := mux.NewRouter()
	hService := healthService.New()
	r.HandleFunc("/healthz", hService.HealthInfo).Methods("GET")
	n := negroni.Classic() // Includes some default middlewares
	// n.Use(negroni.HandlerFunc(auth.BasicAuthMiddleware))
	n.UseHandler(r)
	go http.ListenAndServe(":9000", n)

	// Employee Service on port :8000
	eRepo := employeeRepository.NewInMem()
	eService := employeeService.NewService(eRepo)
	eEndpoints := employeeEndpoints.New(eService)
	eHandler := employeeTransport.NewHTTPServer(eEndpoints)
	http.ListenAndServe(":8000", eHandler)
}
