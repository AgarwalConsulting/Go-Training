package main

import "fmt"

// Iteratively
// 5! => 1 * 2 * 3 * 4 * 5
// func fac(n int) int {
// 	var res int = 1

// 	for i := 1; i <= n; i++ {
// 		res = res * i
// 	}

// 	return res
// }

// Recursive => Exit condition!
// 5! => 4! * 5
// 5! => 3! * 4 * 5
// 5! => 2! * 3 * 4 * 5
// 5! => 1! * 2 * 3 * 4 * 5
// 5! => 0! * 2 * 3 * 4 * 5
// 5! => 1 * 2 * 3 * 4 * 5

var memo = map[int]int{}

// fac(5) => fac(4) * 5
// fac(5) => fac(4) * 5
func fac(n int) int {
	if n <= 1 {
		return 1
	}
	// Memoize
	if val, ok := memo[n]; ok {
		// Returning the value from map itself
		fmt.Println("\tReturning fac of:", n, "from memo as:", val)
		return val
	}

	fmt.Println("\tComputing the fac of:", n)
	val := fac(n-1) * n
	memo[n] = val
	return val

	// return fac(n-1) * n
}

func main() {
	fmt.Println("Factorial of 0:", fac(0)) // Result: 1
	fmt.Println(memo)
	fmt.Println("-----------------")

	fmt.Println("Factorial of 1:", fac(1)) // Result: 1
	fmt.Println(memo)
	fmt.Println("-----------------")

	fmt.Println("Factorial of 2:", fac(2)) // Result: 1 * 2 => 2
	fmt.Println(memo)
	fmt.Println("-----------------")

	fmt.Println("Factorial of 5:", fac(5)) // Result: 1 * 2 * 3 * 4 * 5 => 120
	fmt.Println(memo)
	fmt.Println("-----------------")

	fmt.Println("Factorial of 6:", fac(6)) // Result: 720
	fmt.Println(memo)
	fmt.Println("-----------------")

	fmt.Println("Factorial of 5:", fac(5)) // Result: 1 * 2 * 3 * 4 * 5 => 120
}
