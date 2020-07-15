package main

import (
	"fmt"
	"sync"
	"time"
)

func fibonacci(id int, c chan<- int, quit <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	x, y := 0, 1
	for {
		select {
		case c <- x:
			fmt.Println("Sent next value!", id)
			x, y = y, x+y
		case <-quit:
			fmt.Println("quiting... ", id)
			return
		default:
			fmt.Println("Waiting for receiver...", id)
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	var wg sync.WaitGroup
	wg.Add(2)

	go fibonacci(1, c, quit, &wg) // Producer
	go fibonacci(2, c, quit, &wg) // Producer

	wg.Wait()
}
