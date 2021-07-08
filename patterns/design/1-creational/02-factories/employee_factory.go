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

func NewEmployeeFactory(designation string, salary float64) func(string) *Employee {
	return func(Name string) *Employee {
		e := Employee{Designation: designation, Salary: salary}

		e.Name = Name
		return &e
	}
}

func main() {
	ceoFactory := NewEmployeeFactory("CEO", 9000000)
	developerFactory := NewEmployeeFactory("Developer", 800000)

	ceo := ceoFactory("Gaurav")

	dev1 := developerFactory("GA")
	dev2 := developerFactory("AF")

	fmt.Println(ceo)
	fmt.Println(dev1, dev2)
}
