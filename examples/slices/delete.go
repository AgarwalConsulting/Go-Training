package main

import "fmt"

func describe(label string, x []string) {
	fmt.Printf("\t label: %s -> %#v, Len: %d, Cap: %d\n", label, x, len(x), cap(x))
}

func delete(emp []string, name string) []string {
	newEmp := make([]string, 0, cap(emp))
	// var newEmp []string

	describe("New emp", newEmp)

	for _, el := range emp {
		if name != el {
			newEmp = append(newEmp, el)
		}

		describe("New emp", newEmp)
	}

	return newEmp
}

func main() {
	e := []string{"Gaurav", "Superman", "Batman"}

	describe("Emp", e) // 3, 3

	// e = e[0:2] // Batman has left

	e = delete(e, "Superman")

	describe("Emp", e)
}
