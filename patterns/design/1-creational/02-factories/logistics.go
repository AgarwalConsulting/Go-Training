package main

// Factory Function: Design a NewPerson for Person type
// Interface Factory: Return an interface instead
// Factory Generator: Design an Employee
//	(functional): NewEmployeeFactory(string, int) func(Name string) *Employee
//	(struct): EmployeeFactory#Create
// Prototype Factory: Switch on the type of employee: manager, developer, etc.

type Transport interface {
	deliver()
}

type Truck struct {
}

func (t Truck) deliver() {}

type Ship struct{}

func (s Ship) deliver() {}

type Logistics struct {
	o Order
}

func (l Logistics) createTransport() Transport {

}

type Order struct {
}

func main() {
	// o := Order{Truck{}}
	o := Order{}
	Logistics{o}

	t := logistics.createTransport()

	t.deliver()
}
