package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// func fibonacci(id int, c chan<- int, quit <-chan int, wg *sync.WaitGroup) {
func fibonacci(ctx context.Context, id int, c chan<- int, wg *sync.WaitGroup) {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if id < 10 {
		childCtx, _ := context.WithCancel(ctx)

		go fibonacci(childCtx, 10+id, c, nil)
	}

	x, y := 0, 1
	for {
		select {
		case c <- x: // Sending will block unless there is a receiver
			fmt.Println("Sent next value!", id)
			x, y = y, x+y
		case <-ctx.Done(): // Receive will unblock if the ctx is cancelled
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
	fmt.Printf("%T\n", ctx)

	go func() { // Consumer
		defer fmt.Println("Terminating consumer...")
		for i := 0; i < 10; i++ {
			fmt.Println("Received", i, "th value: ", <-c)
			time.Sleep(1 * time.Second)
		}
		cancelFn()
	}()

	var wg sync.WaitGroup // Counter: 0
	wg.Add(3)             // 2

	go fibonacci(ctx, 1, c, &wg) // Producer
	go fibonacci(ctx, 2, c, &wg) // Producer
	go fibonacci(ctx, 3, c, &wg) // Producer

	wg.Wait() // Blocks until counter reaches 0
}
