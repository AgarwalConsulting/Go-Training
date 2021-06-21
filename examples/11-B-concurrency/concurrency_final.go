package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
	"sync"
)

// For Windows
// func getRandomIterCount() int {
// 	randomNumber := rand.Intn(100)

// 	iterCount := 300 + randomNumber

// 	return iterCount
// }

// For Unix
func getRandomIterCount() int {
	randomFile, _ := os.Open("/dev/random")
	randomReader := bufio.NewReader(randomFile)
	randomNumber, _ := rand.Int(randomReader, big.NewInt(100))
	iterCount := int(700 + randomNumber.Int64())

	return iterCount
}

func busyWork(iterCount int) {
	for i := 0; i < iterCount; i++ {
		for j := 0; j < iterCount; j++ {
			mul := []byte(strconv.Itoa(i * j))       // CPU intensive
			ioutil.WriteFile("/dev/null", mul, 0644) // IO intensive
		}
	}
}

func say(s string, index int, wg *sync.WaitGroup) {
	fmt.Println("From - Say: ", index, "is starting...")
	iterCount := getRandomIterCount()

	// CPU & IO intensive work for random iterations
	busyWork(iterCount)

	fmt.Printf("Hello, %s! from %d say fn; Iter Count: %d\n", s, index, iterCount)
	wg.Done() // Decrements the counter by 1
}

func main() {
	var wg sync.WaitGroup // Counter: 0
	for i := 0; i < 4; i++ {
		wg.Add(1)
		fmt.Println("Started say", i, "...")
		go say("world", i, &wg) // Goroutine => light-weight threads => Run independently
	}

	fmt.Println("Started all say goroutines!")
	wg.Wait() // Blocks until the counter reaches 0 again
	// time.Sleep(30 * time.Second)
	fmt.Println("Exiting main...")
}
