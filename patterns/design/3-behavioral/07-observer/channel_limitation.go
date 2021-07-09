package main

import (
	"fmt"
	"time"
)

func Consumer(id int, c chan int) {
	for x := range c {
		fmt.Println("Received: ", x, "by consumer:", id)
	}
}

func Producer(id int, c chan<- int) {
	i, j := -1, 1

	for {
		i, j = j, j+i
		c <- j
		fmt.Println("Sent: ", j, "by producer:", id)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	c := make(chan int)

	go Producer(1, c)
	// go Producer(2, c)

	go Consumer(1, c)
	go Consumer(2, c)

	for {
	}
}
