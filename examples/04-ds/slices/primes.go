package main

import "fmt"

func describe(s []int) {
	if s == nil {
		fmt.Printf("-> slice is nil\n\t")
	}
	fmt.Println("Len: ", len(s), " Cap: ", cap(s), "Values: ", s)
}

func main() {
	var primes []int // Slices are references to underlying array
	fmt.Printf("%T %v\n", primes, primes)

	if primes == nil {
		fmt.Println("Slices are nil by default!")
	}

	// primes = nil

	primes = []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23}

	// primes = make([]int, 0, 10)
	// primes = make([]int, 10)

	describe(primes)

	// primes = append(primes, 1, 2, 3, 5)

	// describe(primes) // Len: 4; Cap: 10

	// primes = append(primes, 7, 11) // No malloc!

	// describe(primes) // Len: 6; Cap: 10

	// primes = append(primes, 13, 17, 19)

	// describe(primes) // Len: 9; Cap: 10

	// primes = append(primes, 23)

	// describe(primes) // Len: 10; Cap: 10

	// primes = append(primes, 7, 11, 13, 17, 19)

	// describe(primes) // Len: 9; Cap: 12
}
