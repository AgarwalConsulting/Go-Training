package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// type Employee map[string]interface{}
type Employee struct {
	Name        string
	Designation string
	Department  string
	ProjectID   int
}

func main() {
	gopath := os.Getenv("GOPATH")
	f, err := os.Open(gopath + "/src/github.com/AgarwalConsulting/Go-Training/exercises/basic/reading/employees.json")

	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	defer f.Close()

	// var e = Employee{}
	var e []Employee

	json.NewDecoder(f).Decode(&e)

	// fmt.Println("Name:", e["Name"])
	// fmt.Println("Name:", e.Name)

	var projectIDs = map[int]struct{}{}

	for _, emp := range e {
		projectIDs[emp.ProjectID] = struct{}{}
	}

	fmt.Println("Unique Project IDs:")

	for projectID := range projectIDs {
		fmt.Println("\t", projectID)
	}
}
