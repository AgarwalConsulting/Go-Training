package main

import "fmt"

// func IsEqual(x, y int) bool {
// 	return x == y
// }

// func IsEqual(x, y interface{}) bool {
// func IsEqual(x, y any) bool {
// 	return x == y
// }

func IsEqual[T comparable](x, y T) bool {
	return x == y
}

func main() {
	fmt.Println(IsEqual(10, 12))
	fmt.Println(IsEqual("hello", "hello"))
	// fmt.Println(IsEqual("hello", 12))
}
