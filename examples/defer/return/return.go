// https://tour.golang.org/flowcontrol/12
package main

import (
	"fmt"
)

func say() (x string) {
	defer func() {
		fmt.Println("Defer called...")
		x = "hello"
		fmt.Println("Defer exited!")
	}()

	x = "bye"

	// panic("oh no!")
	// os.Exit(0)

	return
}

func main() {
	fmt.Println(say())
}
