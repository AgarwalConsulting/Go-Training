package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int

	go func() {
		fmt.Println("Sending...")
		time.Sleep(1 * time.Second)

		ch <- 42 // Send a value; send blocks until there is a receive

		fmt.Println("Sent!")
	}()

	fmt.Println("Receiving...")
	time.Sleep(1 * time.Second)

	x := <-ch // Receiving a value from the channel; receive blocks until there is a send

	fmt.Println("Received: ", x)
}
