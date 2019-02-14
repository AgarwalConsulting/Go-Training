package main

import (
	"fmt"
	"sync"
)

func say(s string, wg *sync.WaitGroup) {
	for i := 0; i < 500; i++ {
		for j := 0; j < 500; j++ {
			mul := i * j
			fmt.Println(mul)
		}
	}

	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		go say("world", &wg)
		wg.Add(1)
	}

	wg.Wait()
}
