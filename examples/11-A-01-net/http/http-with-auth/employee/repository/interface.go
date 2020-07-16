package repository

import "algogrit.com/emp-server/employee/entities"

// EmployeeRepository represents a employee store
type EmployeeRepository interface {
	RetrieveAll() []entities.Employee
	FindBy(int) *entities.Employee
	Save(entities.Employee) (*entities.Employee, error)
	Update(int, entities.Employee) error
}
