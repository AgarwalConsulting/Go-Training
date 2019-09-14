package main

import "fmt"

func describe(desc string, slice []string) {
	fmt.Printf("\t%s -> len: %d, cap: %d, %s\n", desc, len(slice), cap(slice), slice)
}

func main() {
	names := []string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	describe("names", names)

	a := names[1:3]

	// b := [2]string{"George", "Ringo"}

	describe("a", a)

	fmt.Println("Appending... XXX to a")
	// a[2] = "XXX"
	b := append(a, "XXX")

	describe("names", names)
	describe("a", a)
	describe("b", b)

	fmt.Println("Appending... Tarun to names & Updating last name to Ringo")

	newNames := append(names, "Tarun")
	names[3] = "Ringo"

	describe("names", names)
	describe("newNames", newNames)

	fmt.Println("Resizing a...")
	a = a[0:cap(a)]

	describe("a", a)

	fmt.Println("Appending a slice to another slice...")

	// arr := [5]string{}
	c := make([]string, 0, 5)

	describe("c", c)
	c = c[0:2]
	describe("c", c)

	c[0] = "Gaurav"
	c[1] = "Sachin"

	describe("c", c)

	d := append(c, a...)

	describe("a", a)
	describe("c", c)
	describe("d", d)
}
