package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		for {
			fmt.Println("Receiving...")
			fmt.Println("Received: ", <-ch) // Receive from a channel blocks until there is a send
		}
	}()

	i, j := 0, 1
	for iter := 0; iter < 10; iter++ {
		fmt.Println("Sending...")
		ch <- i // Send is going to block until there is a receive
		fmt.Println("Sent: ", i)
		i, j = j, i+j
	}
}
