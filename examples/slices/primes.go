package main

import "fmt"

func describe(s []int) {
	if s == nil {
		fmt.Printf("-> slice is nil\n\t")
	}
	fmt.Println("Len: ", len(s), " Cap: ", cap(s), "Values: ", s)
}

func main() {
	var primes []int // Len: 0, Cap: 0
	describe(primes)

	// var a = [6]int{2, 3, 5, 7, 11, 13}

	// primes = a[0:4]

	// describe(primes)

	// primes[2] = 17

	// describe(primes)

	// fmt.Println(a) // [2, 3, 17, 7, 11, 13]

	// make([]<type>, len) -> len == cap
	// make([]<type>, len, cap)

	// primes = make([]int, 10)
	// describe(primes)

	primes = make([]int, 0, 5)
	describe(primes)

	// primes = make([]int, 0)
	// describe(primes)

	primes = append(primes, 2, 3)
	describe(primes) // Len: 2, Cap: 5

	primes = append(primes, 5)
	describe(primes) // Len: 3, Cap: 5

	primes = append(primes, 7)
	describe(primes) // Len: 4, Cap: 5

	primes = append(primes, 11)
	describe(primes) // Len: 5, Cap: 5

	primes = append(primes, 13)
	describe(primes) // Len: 6, Cap: 10
}
