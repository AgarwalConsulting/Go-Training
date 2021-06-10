package main

import "fmt"

func main() {
	// var s []rune = []rune{'h', 'e', 'l', 'l', 'o'}

	// fmt.Println(string(s))

	str := "hello 😀"

	var b []byte = []byte(str)
	var r []rune = []rune(str)

	fmt.Println(str[0:7])       // []
	fmt.Println(b[0:7])         // []
	fmt.Println(string(b[0:7])) // []
	fmt.Println(r[0:7])         // []
	fmt.Println(string(r[0:7])) // []
}
