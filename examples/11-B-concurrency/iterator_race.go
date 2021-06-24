package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		// go fmt.Println(i)
		// k := i
		go func(k int) {
			fmt.Println(k)
		}(i)
	}
	time.Sleep(1 * time.Second)
}
