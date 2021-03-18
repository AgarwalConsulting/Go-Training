package main

import "fmt"

func desc(s1, s2 []int) {
	fmt.Printf("s1: %p ; s2: %p\n", s1, s2)
	fmt.Println("\ts1:", s1, "|\ts2:", s2)
}

func main() {
	var s1 = []int{1, 2, 3} // Len: 3

	s2 := make([]int, len(s1))

	desc(s1, s2)

	copy(s2, s1) // dst, src; minimum of (len(s1), len(s2)) == 0

	desc(s1, s2) // s1: [1, 2, 3] ; s2: []
}
