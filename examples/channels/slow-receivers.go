package main

import (
	"fmt"
	"time"
)

func Consumer(c chan int, id int) {
	for {
		fmt.Println("C", id, "Processing...")
		time.Sleep(1 * time.Second)
		i := <-c // Receive blocks if the buffer is empty
		fmt.Println("C", id, "Processed: ", i)
	}
}

func main() {
	var c chan int // nil
	c = make(chan int, 10)

	go Consumer(c, 1)
	go Consumer(c, 2)

	// time.Sleep(3 * time.Second)
	i, j := 0, 1
	for x := 0; x < 2; x++ {
		i, j = j, i+j
		nextVal := j - i
		c <- nextVal // Sending blocks only if the buffer is full
		fmt.Println("Sent: ", nextVal)
	}
}
