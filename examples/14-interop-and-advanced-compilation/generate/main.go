// More info: https://blog.golang.org/generate
package main

import "fmt"

//go:generate go run gen.go

func main() {
	fmt.Println(strFunc())
}
