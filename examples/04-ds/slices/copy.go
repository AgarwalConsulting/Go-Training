package main

import "fmt"

func main() {
	var s1 = []int{1, 2, 3} // Len: 3
	// var s2 = []int{6, 7}    // Len: 2
	s2 := make([]int, len(s1))

	// copy(s1, s2)
	copy(s2, s1)

	fmt.Println(s1) // [1, 2, 3]
	fmt.Println(s2) // [1, 2]
}
