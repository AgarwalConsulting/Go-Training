package main

import "fmt"

func main() {
	// var s []int --- nil slice
	s := make([]int, 5) // len and cap as 5
	// s := make([]int, 1, 5) // len as 1 and cap as 5

	fmt.Println(s, len(s), cap(s))

	s = s[0:3]

	fmt.Println(s, len(s), cap(s))
}
