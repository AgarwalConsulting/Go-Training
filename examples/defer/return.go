// https://tour.golang.org/flowcontrol/12
package main

import (
	"fmt"
)

func say(n int) (x string) {
	defer func() {
		fmt.Println("Defer called...")
		fmt.Println("Value of x: ", x)
		x = "hello"
		fmt.Println("Defer exited!")
	}()

	x = "bye"

	if n == 42 {
		return "The answer to the life, the universe and everything!"
	}

	return x
}

func main() {
	fmt.Println(say(0))

	fmt.Println(say(42))
}
