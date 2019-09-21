package main

import (
	"context"
	"fmt"
	"time"
)

func printHello(ctx context.Context) {
	time.Sleep(1 * time.Second)
	fmt.Println("Does it print hello?")

	select {
	case <-time.After(250 * time.Microsecond):
		fmt.Println("Hello")
	case <-ctx.Done():
		fmt.Println("Returning from printHello!")
	}
}

func fibonacci(fibCtx context.Context, c chan int) {
	x, y := 0, 1

	helloCtx, cancelFn := context.WithCancel(fibCtx)
	defer cancelFn()

	go printHello(helloCtx)

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

	fibCtx, _ := context.WithTimeout(ctx, 200*time.Microsecond)

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()
	fibonacci(fibCtx, c)

	time.Sleep(2 * time.Second)
}
