package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"sync"
)

func say(s string, index int) {
	randomFile, _ := os.Open("/dev/random")
	randomReader := bufio.NewReader(randomFile)
	randomNumber, _ := rand.Int(randomReader, big.NewInt(5))
	iterCount := int(10 + randomNumber.Int64())

	for i := 0; i < iterCount; i++ {
		for j := 0; j < iterCount; j++ {
			mul := strconv.Itoa(i * j)
			message := fmt.Sprintf("From %d goroutine - printing: %s \n", index+1, mul)

			f, err := os.OpenFile("/tmp/mul.txt", os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				panic(err)
			}

			if _, err = f.WriteString(message); err != nil {
				panic(err)
			}

			f.Close()
		}
	}

	fmt.Printf("Hello, %s! from %d goroutine; Iter Count: %d\n", s, index+1, 10)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		j := i
		wg.Add(1)
		go func() {
			say("world", j)
			wg.Done()
		}()
	}

	wg.Wait()

	// time.Sleep(1 * time.Second)
}
