package main

import (
	"fmt"
	"sync"
	"time"
)

func Producer(c chan<- int) {
	// time.Sleep(3 * time.Second)
	i, j := 0, 1
	for {
		i, j = j, i+j
		nextVal := j - i
		c <- nextVal // Sending blocks only if the buffer is full
		fmt.Println("Sent: ", nextVal)
	}
}

func Consumer(c <-chan int, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for x := 0; x < 2; x++ {
		fmt.Println("C", id, "Processing...")
		i := <-c // Receive blocks if the buffer is empty
		// Take some time to process that value
		time.Sleep(1 * time.Second)
		fmt.Println("C", id, "Processed: ", i)
	}
}

func main() {
	var c chan int

	fmt.Println(c) // <nil>

	c = make(chan int)

	go Producer(c)

	var wg sync.WaitGroup
	wg.Add(2)

	go Consumer(c, 2, &wg)
	go Consumer(c, 1, &wg)

	wg.Wait()
}
