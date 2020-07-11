package main

type Employee struct {
	Person
	Designation      string
	YearOfExperience int
}

// func (e Employee) String() string {
// 	// return "Name: " + e.Name + " | Age: " + strconv.Itoa(e.Age) + " | Designation: " + e.Designation + " | Experience: " + strconv.Itoa(e.YearOfExperience)
// 	return e.Person.String() + " | Designation: " + e.Designation + " | Experience: " + strconv.Itoa(e.YearOfExperience)
// }
