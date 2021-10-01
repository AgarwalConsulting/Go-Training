package main

import "fmt"

func main() {
	// var s []rune = []rune{'h', 'e', 'l', 'l', 'o'}

	// fmt.Println(string(s))

	str := "hello 😀"

	fmt.Println("Length:", len(str))

	var b []byte = []byte(str)

	fmt.Println(str[0:7])       //
	fmt.Println(b[0:7])         //
	fmt.Println(string(b[0:7])) //

	// var r []rune = []rune(str)

	// for i := 0; i < len(r); i++ {
	// 	char := r[i]
	// 	fmt.Println("\tChar at index", i, ":", char, "as string:", string(char))
	// }

	// fmt.Println(r)
	// fmt.Println(r[0:7])         //
	// fmt.Println(string(r[0:7])) //
}
