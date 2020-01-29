package main

import (
	"context"
	"fmt"
	"time"
)

func fibonacci(fibCtx context.Context, c chan int) {
	x, y := 0, 1

	// helloCtx, cancelFn := context.WithCancel(fibCtx)
	// defer cancelFn()

	for {
		select {
		case c <- x: // send
			x, y = y, x+y
		case <-fibCtx.Done(): // receive
			fmt.Println("Returning from fibonacci...")
			return
		}
	}
}

func main() {
	c := make(chan int)

	ctx := context.Background()

	fibCtx, _ := context.WithTimeout(ctx, 100*time.Microsecond)

	// Consumer
	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	fibonacci(fibCtx, c)

	time.Sleep(2 * time.Second)
}
