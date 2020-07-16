package service

import (
	"context"

	"algogrit.com/emp-server/employee/entities"
	"algogrit.com/emp-server/employee/repository"
)

type empService struct {
	repo repository.EmployeeRepository
}

func (es *empService) Show(ctx context.Context, ID int) (entities.Employee, error) {
	employee := es.repo.FindBy(ID)

	return *employee, nil
}

func (es *empService) Create(ctx context.Context, newEmployee entities.Employee) (entities.Employee, error) {
	employee, err := es.repo.Save(newEmployee)
	return *employee, err
}

func NewService(repo repository.EmployeeRepository) EmployeeService {
	return &empService{repo}
}
