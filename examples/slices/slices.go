package main

import "fmt"

func main() {
	names := []string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]

	a = append(a, "Hi")
	a = append(a, "Hi1")

	b := names[1:3]
	b = append(b, "Hi2")

	b = append(b, "Hi3")

	fmt.Println(a, b)

	a = append(a, "Hi4")

	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
