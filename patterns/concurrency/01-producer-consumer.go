package main

import (
	"context"
	"fmt"
)

// Producer
func fibonacci(ctx context.Context, ch chan<- int) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-ctx.Done():
			fmt.Println("Stopping Producer!")
			return
		}
	}
}

func main() {
	c := make(chan int)
	ctx := context.Background()

	cancelableCtx, cancelFn := context.WithCancel(ctx)

	// Consumer
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println(<-c)
		}
		fmt.Println("Cancelling context...")
		cancelFn()
	}()

	go fibonacci(cancelableCtx, c)
	go fibonacci(cancelableCtx, c)

	for {

	}
	// <-quit
}
