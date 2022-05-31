package main

import (
	"fmt"
	"math"
)

func collect[T any, U any](in []T, op func(T) U) []U {
	// var out []U
	out := make([]U, 0, len(in))

	for _, el := range in {
		out = append(out, op(el))
	}

	return out
}

func main() {
	primes := []int{1, 2, 3, 5, 7}

	fmt.Println(primes)

	// Get double of all primes

	// var doubled []float64

	// for _, el := range primes {
	// 	doubled = append(doubled, el*2)
	// }

	doubled := collect(primes, func(el int) int { return el * 2 })

	fmt.Println("Doubled Primes:", doubled, len(doubled), cap(doubled))

	// Get sqrt of all primes

	sqrt := collect(primes, func(el int) float64 { return math.Sqrt(float64(el)) })
	// sqrt := collect(primes, math.Sqrt)

	fmt.Println("Sqrt:", sqrt)
}
