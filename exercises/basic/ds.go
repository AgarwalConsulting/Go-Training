package main

import "fmt"

func main() {
	for {
		fmt.Println("1. Add person to department")
		fmt.Println("2. Show people in department")
		fmt.Println("3. Show all people in company")
		fmt.Println("Enter your choice:")
		var choice string

		// Retrieve choice from user and switch...
		// ?

		switch choice {
		case "1":
			// Prompt user for person's name & department and then store it
			// addPersonToDepartment()
		case "2":
			// Prompt user for department and then display people in a department
			// showPeopleInDeparment()
		case "3":
			// Display all the people sorted across departments
			// showPeopleInCompany()
		}
	}
}
