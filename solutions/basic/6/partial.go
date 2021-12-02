package main

import (
	"fmt"
	"strings"
)

type Company map[string][]string

func addPersonToDepartment(company map[string][]string, department string, person string) {
	company[department] = append(company[department], person)
}

func showPeopleInDeparment(company map[string][]string, department string) {
	peopleInDepartment := company[department] // Type: []string

	fmt.Println("-------------------------------------")

	fmt.Println("Department:", department, "has:")

	for _, person := range peopleInDepartment {
		fmt.Println("\t", person)
	}
	fmt.Println("-------------------------------------")
}

func main() {
	var company = Company{}

	for {
		fmt.Println("===================================================")
		fmt.Println("1. Add person to department")
		fmt.Println("2. Show people in department")
		fmt.Println("3. Show all people in company")
		fmt.Println("4. Quit")
		fmt.Print("Enter your choice:\t")
		var choice string
		// Retrieve choice from user and switch...
		fmt.Scanln(&choice)
		fmt.Println("===================================================")

		choice = strings.Trim(choice, "\n\t")

		switch choice {
		case "1":

			// Prompt user for person's name & department and then store it
			fmt.Print("Enter the name and department:\t")

			var person, department string
			fmt.Scanf("Add %s to %s", &person, &department) // Add <person-name> to <department>

			addPersonToDepartment(company, department, person)
		case "2":
			var department string
			// Prompt user for department and then display people in a department
			fmt.Print("Enter the department:\t")
			fmt.Scanf("%s", &department) // Add <person-name> to <department>

			showPeopleInDeparment(company, department)
		case "3":
			// Display all the people sorted across departments
			// showPeopleInCompany()
		default:
			return
		}
	}
}
