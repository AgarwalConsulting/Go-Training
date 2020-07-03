package service

import "net/http"

type EmployeeService interface {
	Index(http.ResponseWriter, *http.Request)
	Show(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
}
