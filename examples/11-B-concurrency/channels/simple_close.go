package main

import (
	"fmt"
	"time"
)

func fibGenerator(c chan<- int) {
	i, j := 0, 1
	for x := 0; x < 20; x++ {
		i, j = j, i+j // 1, 0+1
		nextVal := j - i
		c <- nextVal // Sending blocks only if the buffer is full
		fmt.Println("Sent: ", nextVal)
	}
	fmt.Println("Completed generator!")
	// close(c)
}

func Consumer(c <-chan int, id int) {
	for {
		fmt.Println("C", id, "Processing...")
		time.Sleep(100 * time.Millisecond)
		i := <-c // Receive blocks if the buffer is empty
		fmt.Println("C", id, "Processed: ", i)
	}
}

func main() {
	var c chan int // nil
	c = make(chan int, 10)

	go fibGenerator(c)

	Consumer(c, 1)
}
