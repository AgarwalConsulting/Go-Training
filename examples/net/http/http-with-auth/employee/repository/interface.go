package repository

// EmployeeRepository represents a employee store
type EmployeeRepository interface {
	RetrieveAll() []Employee
	FindBy(int) *Employee
	Save(Employee) (*Employee, error)
	Update(int, Employee) error
}
