package service

import (
	"context"

	"algogrit.com/emp-server/employee/entities"
)

type EmployeeService interface {
	Show(ctx context.Context, ID int) (entities.Employee, error)
	Create(ctx context.Context, newEmployee entities.Employee) (entities.Employee, error)
}
