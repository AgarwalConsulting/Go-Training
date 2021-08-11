package main

import (
	"fmt"
	"sync"
	"time"
)

func Producer(ch chan<- int) {
	i, j := -1, 1

	for {
		i, j = j, j+i

		fmt.Println("Sending...")
		ch <- j // Sending the next value into the channel
		fmt.Println("Sent:", j)
	}
}

func Consumer(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	var x = 0
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(x) * time.Second)
		fmt.Println("Waiting to receive...", id)
		x = <-ch
		fmt.Printf("Received in C %d: %d\n", id, x)
	}
}

func main() {
	ch := make(chan int)

	go Producer(ch)

	var wg sync.WaitGroup

	go Consumer(1, ch, &wg)
	go Consumer(2, ch, &wg)

	wg.Add(2)

	wg.Wait()
}
