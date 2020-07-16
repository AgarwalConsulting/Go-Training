package endpoints

import (
	"algogrit.com/emp-server/employee/service"
	"github.com/go-kit/kit/endpoint"
)

type Endpoint struct {
	Show   endpoint.Endpoint
	Create endpoint.Endpoint
}

func New(s service.EmployeeService) Endpoint {
	var e Endpoint

	e.Show = makeShowEndpoint(s)
	e.Create = makeCreateEndpoint(s)

	return e
}
