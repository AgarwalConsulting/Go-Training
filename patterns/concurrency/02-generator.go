package main

import "fmt"

func fib(n int) chan int {
	fibChan := make(chan int)
	go func() {
		for i, j := 0, 1; i < n; i, j = i+j, i {
			fibChan <- i
		}
		close(fibChan)
	}()
	return fibChan
}

func main() {
	for i := range fib(1000) {
		fmt.Println(i)
	}
}
