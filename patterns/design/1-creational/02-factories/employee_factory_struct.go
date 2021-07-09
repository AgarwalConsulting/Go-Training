package main

import "fmt"

// Factory Function: Design a NewPerson for Person type
// Interface Factory: Return an interface instead
// Factory Generator: Design an Employee
//	(functional): NewEmployeeFactory(string, int) func(Name string) *Employee
//	(struct): EmployeeFactory#Create
// Prototype Factory: Switch on the type of employee: manager, developer, etc.

type Employee struct {
	Name        string
	Designation string
	Salary      float64
}

func NewEmployee(name string, kind string) Employee {
	switch kind {
	case "CEO":
		return Employee{name, "CEO", 9_000_000}
	case "Developer":
		return Employee{name, "Developer", 800_000}
	}
	return Employee{name, kind, 0}
}

type EmployeeFactory struct {
	designation string
	salary      float64
}

func (ef EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, ef.designation, ef.salary}
}

func main() {
	ceoFactory := EmployeeFactory{"CEO", 9_000_000}
	developerFactory := EmployeeFactory{"Developer", 800_000}

	ceo := ceoFactory.Create("Gaurav")

	dev1 := developerFactory.Create("GA")
	dev2 := developerFactory.Create("AF")

	fmt.Println(ceo)
	fmt.Println(dev1, dev2)
}
