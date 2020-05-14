package main

import (
	"fmt"
	"time"
)

func Consumer(c <-chan int) {
	for {
		fmt.Println("Processing...")
		time.Sleep(2 * time.Second)
		i := <-c
		fmt.Println("Processed: ", i)
	}
}

func main() {
	c := make(chan int)

	go Consumer(c)

	i, j := 0, 1
	for {
		i, j := j, i+j
		fmt.Println("Sending...")
		c <- j - i
	}
}
