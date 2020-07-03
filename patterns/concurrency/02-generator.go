package main

import (
	"fmt"
)

func fib(n int) chan int {
	fibChan := make(chan int, 100)
	go func() {
		for i, j := 0, 1; i < n; i, j = i+j, i {
			fibChan <- i
			fmt.Println("Producing: ", i)
		}
		close(fibChan)
		fmt.Println("Closed channel!")
	}()
	return fibChan
}

func main() {
	fibChan := fib(1000)

	fmt.Println(<-fibChan)
	fmt.Println(<-fibChan)
	// for i := range fib(1000) {
	// 	fmt.Println(i)
	// 	time.Sleep(time.Millisecond * 100)
	// }
}
