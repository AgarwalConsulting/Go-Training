package main

import "fmt"

func desc(s1 []int) {
	fmt.Printf("s1: %p; Len: %d, Cap: %d; %#v\n", s1, len(s1), cap(s1), s1)

	for idx, el := range s1 {
		fmt.Println("\t", idx, ":", el)
	}
}

func main() {
	// s := make([]int, 5)
	s := []int{1, 2, 3, 5, 7}

	desc(s)

	s = s[0:1]

	desc(s)

	s = s[0:cap(s)]

	desc(s)
}
