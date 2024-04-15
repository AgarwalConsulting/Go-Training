package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// func fibonacci(id int, c chan<- int, quit <-chan int, wg *sync.WaitGroup) {
func fibonacci(ctx context.Context, id int, c chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	if id < 10 {
		childCtx, _ := context.WithCancel(ctx)
		wg.Add(1)
		go fibonacci(childCtx, id+10, c, wg)
	}
	x, y := 0, 1
	for { // Forever!
		select {
		case c <- x: // Sending to the data channel
			fmt.Println("Sent next value!", id)
			x, y = y, x+y
		case <-ctx.Done(): // Receiving from a closed channel
			fmt.Println("!!!quiting...!!! ->", id)
			return
		default:
			fmt.Println("---Waiting for receiver...", id)
			time.Sleep(time.Millisecond * 900)
		}
	}
}

func main() {
	c := make(chan int) // Data channel // Unbuffered

	ctx, cancelFn := context.WithCancel(context.Background())

	// quit := make(chan int) // Signal channel
	go func() { // Consumer
		defer fmt.Println("Terminating consumer...")
		for i := 0; i < 10; i++ {
			fmt.Println("Received", i, "th value: ", <-c)
			time.Sleep(1 * time.Second)
		}
		// close(quit)
		cancelFn()
	}()

	var wg sync.WaitGroup
	wg.Add(3)

	go fibonacci(ctx, 1, c, &wg) // Producer
	go fibonacci(ctx, 2, c, &wg) // Producer
	go fibonacci(ctx, 3, c, &wg) // Producer

	wg.Wait()
}
