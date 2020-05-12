package main

import "fmt"

func describe(desc string, slice []string) {
	fmt.Printf("\t%s -> len: %d, cap: %d, %s\n", desc, len(slice), cap(slice), slice)
}

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println("names", names)

	a := names[0:2] // []string{"John", "Paul"}
	b := names[1:3] // []string{"Paul", "George"}
	describe("a", a)
	describe("b", b)

	// b[0] = "XXX"
	// describe("a", a)
	// describe("b", b)

	// // a = a[0:3]
	// // a[2] = "Gaurav"
	// a = append(a, "Gaurav")
	// fmt.Println("Resizing a and appending \"Gaurav\"")
	// describe("a", a)
	// describe("b", b)

	fmt.Println("names", names)
}
