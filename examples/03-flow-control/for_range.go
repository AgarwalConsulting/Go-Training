package main

import "fmt"

func main() {
	str := "hello 😀"

	// // Iterating on each byte of a string
	// var char byte // unit8 = 8 bits
	// for i := 0; i < len(str); i++ {
	// 	char = str[i]
	// 	fmt.Printf("%T %#v\n", char, string(char))
	// }

	// // Iterating on each rune of a string
	// // <for index, element := range str>
	// var char rune // int32 = 32 bits or 4 bytes
	// for index, char := range str {
	// 	fmt.Printf("%d: %T %#v\n", index, char, string(char))
	// }

	for _, char := range str {
		fmt.Printf("%T %#v %#v\n", char, char, string(char))
	}
}

// Range works on:
// - string
// - arrays
// - slices
// - maps
// - channels
