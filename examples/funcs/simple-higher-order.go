package main

import (
	"fmt"
	"math"
)

func myMap(nums []float64, fn func(float64) float64) []float64 {
	collector := []float64{}
	for i := 0; i < len(nums); i++ {
		collector = append(collector, fn(nums[i]))
	}

	return collector
}

func main() {
	nums := []float64{1, 2, 3, 4, 5}

	// newNums := myMap(nums, math.Sqrt)
	newNums := myMap(nums, math.Sin)

	fmt.Println(newNums)

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i] * 2)
	// }

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println((nums[i] + 10) * 3)
	// }
}
