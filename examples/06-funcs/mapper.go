package main

import (
	"fmt"
	"math"
)

func collect(nums []float64, operation func(float64) float64) []float64 {
	results := make([]float64, 0, len(nums))

	for _, el := range nums {
		results = append(results, operation(el))
	}

	return results
}

func main() {
	primes := []float64{1, 2, 3, 5, 7}

	doubled := collect(primes, func(el float64) float64 { return el * 2 })

	squared := collect(primes, func(el float64) float64 { return math.Pow(el, 2) })

	fmt.Println(doubled)
	fmt.Println(squared)
}
