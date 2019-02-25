package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

// func say(s string, index int, wg *sync.WaitGroup) {
func say(s string, index int) {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			mul := []byte(strconv.Itoa(i * j))
			ioutil.WriteFile("/dev/null", mul, 0644)
		}
	}

	fmt.Printf("Hello, %s! from %d goroutine\n", s, index+1)

	// wg.Done()
}

func main() {
	// var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		say("world", i)
		// go say("world", i, &wg)
		// wg.Add(1)
	}

	// wg.Wait()
}
