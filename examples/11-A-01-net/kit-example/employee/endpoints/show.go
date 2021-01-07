package endpoints

import (
	"context"

	"algogrit.com/emp-server/employee/entities"
	"algogrit.com/emp-server/employee/service"

	"github.com/go-kit/kit/endpoint"
)

type ShowRequest struct{
	ID int
}

type ShowResponse struct {
	Employee entities.Employee
}

func makeShowEndpoint(s service.EmployeeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		showReq := request.(ShowRequest)

		employee, err := s.Show(ctx, showReq.ID)

		return ShowResponse{Employee: employee}, err
	}
}
