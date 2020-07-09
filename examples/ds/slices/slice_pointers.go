package main

import "fmt"

func main() {
	var s = []int{1, 2, 3} // cap: 3

	// p := &s

	// *p = append(*p, 42)

	fmt.Printf("Before appending: %p\n", &s)
	fmt.Println("s: ", s, len(s), cap(s))

	s = append(s, 42)

	fmt.Printf("After appending: %p\n", &s)

	fmt.Println("s: ", s, len(s), cap(s)) // [1, 2, 3, 42], 4, 6

	_ = append(s, 12)

	fmt.Println("s: ", s, len(s), cap(s))

	s = s[:cap(s)] // Resizing the slice

	fmt.Println("s: ", s, len(s), cap(s))
}
