package repository

type employeeInMem struct {
	idSequence int
	employees  []Employee
}

func (er *employeeInMem) nextID() int {
	er.idSequence++
	return er.idSequence
}

// RetrieveAll fetches all employees from in mem repo
func (er *employeeInMem) RetrieveAll() []Employee {
	return er.employees
}

// FindBy returns an employee with matching id
func (er *employeeInMem) FindBy(id int) *Employee {
	for _, emp := range er.employees {
		if emp.ID == id {
			return &emp
		}
	}

	return nil
}

// Save can save an employee in memory
func (er *employeeInMem) Save(newEmployee Employee) (*Employee, error) {
	newEmployee.ID = er.nextID()

	er.employees = append(er.employees, newEmployee)

	return &newEmployee, nil
}

// Update can update an employee in memory
func (er *employeeInMem) Update(empID int, e Employee) error {
	// TODO: Implement Update

	return nil
}

// NewInMem returns a new in memory employee repository
func NewInMem() EmployeeRepository {
	er := employeeInMem{}

	er.employees = []Employee{
		{er.nextID(), "Algogrit", 1002, "Sr. Java Dev"},
		{er.nextID(), "G A", 1004, "Sr. Go Dev"},
	}

	return &er
}
