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

	a := names[0:2] // []int{"John", "Paul"}
	b := names[1:3] // []int{"Paul", "George"}
	describe("a", a)
	describe("b", b)

	b[0] = "XXX"
	describe("a", a)
	describe("b", b)

	// a = a[0:3]
	// a[2] = "Gaurav"
	a = append(a, "Gaurav")
	fmt.Println("Resizing a and appending \"Gaurav\"")
	describe("a", a)
	describe("b", b)

	b = b[:cap(b)]
	b[2] = "Arjun"
	fmt.Println("Resizing b to cap and appending \"Arjun\"")
	describe("a", a)
	describe("b", b)

	b = append(b, "Srikant")
	fmt.Println("Appending Srikant beyond cap of b")
	describe("a", a)
	describe("b", b)
	fmt.Println("names", names)

	b[0] = "Sunil"
	fmt.Println("Updating b[0] to Sunil")
	describe("a", a)
	describe("b", b)

	b = append(b, "Mary")
	fmt.Println("Appending Mary to b")
	describe("b", b)

	b = append(b, "Vishwas")
	fmt.Println("Appending Vishwas to b")
	describe("b", b)

	b = append(b, "Ashok")
	fmt.Println("Appending Ashok to b")
	describe("b", b)

	fmt.Println("names", names)
}
