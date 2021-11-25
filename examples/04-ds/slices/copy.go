package main

import "fmt"

func desc(s1, s2 []int) {
	fmt.Printf("s1: %p ; s2: %p\n", s1, s2)
	fmt.Println("\ts1:", s1, "|\ts2:", s2)
}

func main() {
	var s1 = []int{1, 2, 3} // Len: 3 | Cap: 3

	var s2 []int // Len: 0 | Cap: 0 // nil!

	desc(s1, s2)

	noOfElements := copy(s2, s1)

	fmt.Println("Copied:", noOfElements, "elements")

	desc(s1, s2)

	s2 = make([]int, 0, 5) // Len: 0 | Cap: 5

	desc(s1, s2)

	noOfElements = copy(s2, s1) // min(len(dest), len(src)) => min(0, 3) => 0

	fmt.Println("Copied:", noOfElements, "elements")

	desc(s1, s2)

	s2 = make([]int, 5) // Len: 5 | Cap: 5

	desc(s1, s2)

	noOfElements = copy(s2, s1) // min(len(dest), len(src)) => min(5, 3) => 3

	fmt.Println("Copied:", noOfElements, "elements")
	desc(s1, s2) // s2 => [1 2 3 0 0]
}
