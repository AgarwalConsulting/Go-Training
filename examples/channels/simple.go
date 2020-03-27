package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 2 // remains blocked until there is a receive
		fmt.Println("Sent: 2")

		c <- 3
		fmt.Println("Sent: 3")

		c <- 42
		fmt.Println("Sent: 42")

		c <- 45
		fmt.Println("Sent: 45")

		c <- 47
		fmt.Println("Sent: 47")
	}()

	for {
		time.Sleep(time.Second * 3)

		val := <-c // remains blocked until there is a send

		fmt.Println("Received:", val)
	}
}
