package main

import "fmt"

func main() {
	str := "hello" // Type: string

	// str := "hello 😀"
	// str := "नमस्ते"

	fmt.Println("Length:", len(str)) // 5 bytes

	var b []byte = []byte(str)

	fmt.Println(b)

	fmt.Println(str[0:5])
	fmt.Println(b[0:5])
	fmt.Println(string(b[0:5]))

	for idx, char := range str {
		fmt.Printf("\tChar at %d: %s (%d)", idx, string(char), char)
	}
}
