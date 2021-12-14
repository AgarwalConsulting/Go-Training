package main

import "fmt"

func main() {
	var a = []int{1, 2}

	var b = []int{3, 4, 5}

	// How to concat both slices?

	var c = make([]int, len(a)+len(b))

	// fmt.Println(a + b)

	copy(c[0:len(a)], a)
	copy(c[len(a):], b)
	// idx := 0
	// for _, el := range a {
	// 	c[idx] = el
	// 	idx++
	// }

	// for _, el := range b {
	// 	c[idx] = el
	// 	idx++
	// }

	fmt.Println(c)
}
