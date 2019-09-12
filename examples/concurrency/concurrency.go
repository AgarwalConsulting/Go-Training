package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
)

func busyWork(iterCount int) {
	for i := 0; i < iterCount; i++ {
		for j := 0; j < iterCount; j++ {
			mul := []byte(strconv.Itoa(i * j))
			ioutil.WriteFile("/dev/null", mul, 0644)
		}
	}
}

func getRandomIterCount() int {
	randomFile, _ := os.Open("/dev/random")
	randomReader := bufio.NewReader(randomFile)
	randomNumber, _ := rand.Int(randomReader, big.NewInt(400))
	iterCount := int(100 + randomNumber.Int64())

	return iterCount
}

func say(s string, index int) {
	iterCount := getRandomIterCount()

	busyWork(iterCount)

	fmt.Printf("Hello, %s! from %d goroutine; Iter Count: %d\n", s, index+1, iterCount)
}

func main() {
	for i := 0; i < 4; i++ {
		say("world", i)
	}

	// time.Sleep(10 * time.Second)
}
