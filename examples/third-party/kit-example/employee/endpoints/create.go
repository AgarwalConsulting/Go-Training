package endpoints

import (
	"context"

	"algogrit.com/emp-server/employee/entities"
	"algogrit.com/emp-server/employee/service"
	"github.com/go-kit/kit/endpoint"
)

type CreateRequest struct {
	Employee entities.Employee
}

type CreateResponse struct {
	Employee entities.Employee
}

func makeCreateEndpoint(s service.EmployeeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		createReq := request.(CreateRequest)

		emp, err := s.Create(ctx, createReq.Employee)

		return CreateResponse{emp}, err
	}
}
