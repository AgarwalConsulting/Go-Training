package main

import "fmt"

func main() {
	str := "नमस्ते" // Na-Ma-S-Te (Grapheme Clustering)
	// str := "hello 😀"
	// str := "hello" // Type: string

	fmt.Println("Length:", len(str)) // 18 bytes

	// fmt.Println(str[0:7]) // "hello 😀" or "hello "

	// fmt.Printf("%d\n", '😀')

	var b []byte = []byte(str) // byte -> uint8 (1 byte -> 8 bits)

	// How much memory would b take: 10 bytes

	fmt.Println(b)

	fmt.Println(len(b)) // Length of b: 10

	// fmt.Println(b[0:7])
	// fmt.Println(string(b[0:7]))

	var r []rune = []rune(str) // rune -> int32 (4 bytes) // Unicode-code point

	fmt.Println("No. of characters:", len(r)) // 6

	fmt.Println(r)

	// fmt.Println(r[0:7])
	// fmt.Println(string(r[0:7]))

	for idx, el := range r {
		fmt.Printf("\t%d: %d - %s\n", idx, el, string(el))
	}
}
