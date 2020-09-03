package main

import "fmt"

func updateSlice(someSlice []int) {
	// someSlice[0] = 11
	someSlice = append(someSlice, 11)
}

func main() {
	primes := [5]int{1, 2, 3, 5, 7}

	s := primes[:]

	fmt.Println("Len:", len(s), "Cap:", cap(s)) // Len: 5, Cap: 5
	fmt.Println(s)                              // [1, 2, 3, 5, 7]

	updateSlice(s)

	fmt.Println("Len:", len(s), "Cap:", cap(s))
	fmt.Println(s)

	fmt.Println(primes)
}
