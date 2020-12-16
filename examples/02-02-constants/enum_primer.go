package main

import (
	"fmt"
)

type status string

const running status = "running"

func main() {
	var s status

	s = running

	fmt.Println(s) // Default value

	fmt.Printf("%v, %T\n", s, s)
}
