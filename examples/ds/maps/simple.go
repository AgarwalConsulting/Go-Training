package main

import "fmt"

type Address struct {
	DoorNo     int
	StreetName string
	Area       string
	City       string
}

func main() {
	var m map[string]Address

	m["Gaurav"] = Address{1, "Ghomti Gali", "Dariya", "Jheel"}

	fmt.Println(m)
}
