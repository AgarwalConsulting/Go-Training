package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
)

func say(s string, index int) {
	iterCount := 100 + rand.Intn(100)

	for i := 0; i < iterCount; i++ {
		for j := 0; j < iterCount; j++ {
			mul := []byte(strconv.Itoa(i * j))
			ioutil.WriteFile("/dev/null", mul, 0644)
		}
	}

	fmt.Printf("Hello, %s! from %d goroutine; Iter Count: %d\n", s, index+1, iterCount)
}

func main() {
	for i := 0; i < 4; i++ {
		say("world", i)
	}

	// time.Sleep(600 * time.Millisecond)
}
