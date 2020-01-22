// https://tour.golang.org/flowcontrol/12
package main

import (
	"fmt"
)

func say(n int) (x string) {
	defer func() {
		fmt.Println("Defer called...")
		fmt.Println("Value of x: ", x)
		// x += "hello"
		// x = "hello"
		fmt.Println("Defer exited!")
	}()

	x = "bye"

	panic("oh no!")
	// os.Exit(0)

	if n == 42 {
		return "The answer to the life, the universe and everything!"
	}

	return x
}

func main() {
	fmt.Println(say(42))
}
