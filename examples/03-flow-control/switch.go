package main

import (
	"fmt"
	"runtime"
)

func main() {
	// fmt.Printf("%T\n", os)

	// implicit `break`; explicit `fallthrough`

	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("Steve rocks!")
	case "linux":
		fmt.Println("Linus rocks!")
	default:
		fmt.Println("This OS sucks!")
	}

	// fmt.Printf("%T\n", os)

	// if os == "linux" {
	// 	fmt.Println("Linus rocks!")
	// } else if os == "darwin" {
	// 	fmt.Println("Steve rocks!")
	// } else {
	// 	fmt.Println("This OS sucks!")
	// }
}
