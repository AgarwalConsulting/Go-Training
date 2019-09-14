// +build !debug

package main

import (
	"fmt"
)

func main() {
	define(DEBUG)
	ifdef(`DEBUG',
		fmt.Println("Debug message...");
	)

	fmt.Println("Release build")
}
