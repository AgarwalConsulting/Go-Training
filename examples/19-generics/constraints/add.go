package main

import "fmt"

// func Add(x, y float64) float64 {
// 	return x + y
// }

// type MyInt int

type Number interface {
	int64 | float64 | uint | uint64 | int
}

func Add[T Number](x, y T) T {
	return x + y
}

// func Add[T float64 | int, U float64 | int](x T, y U) float64 {
// 	return float64(x) + float64(y)
// }

func main() {
	fmt.Println(Add(12.12, 10.12))
	fmt.Println(Add(12, 10))
	// fmt.Println(Add(12.10, 10)) // Doesn't work!
	fmt.Println(Add[float64](12.10, 10))

	// m := MyInt(10)

	// fmt.Println(Add(m, m))

	// var x float64 = 42
	// var y int = 10

	// fmt.Println(Add(x, y))
}
