package extras

import "fmt"

// Hello is a string
var Hello = "hello"

// Hello := "hello"

func init() {
	fmt.Println("Initializing extras...")
	Hello = "Hi"
}
