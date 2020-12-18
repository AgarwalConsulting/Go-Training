package main

import (
	"fmt"
	"time"
)

func main() {
	var x int

	go func() {
		fmt.Println("Sending...")
		time.Sleep(1 * time.Second)

		x = 42
		fmt.Println("Sent")
	}()

	fmt.Println("Receiving...")
	// time.Sleep(1 * time.Second)

	fmt.Println("Received: ", x) // ?
}
