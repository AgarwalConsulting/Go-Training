package main

import (
	"fmt"
)

func main() {
	var ch chan int
	ch = make(chan int) // Unbuffered channel

	go func() {
		fmt.Println("Sending...")
		ch <- 42 // Sending, is blocked until there is a receive
		fmt.Println("Sent!")
	}()

	fmt.Println("Receiving...")
	val := <-ch // Receiving, is blocked until there is a send
	fmt.Println("Received: ", val)
}
