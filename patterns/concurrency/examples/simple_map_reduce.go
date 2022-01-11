package main

import (
	"fmt"
	"sync"
	"time"
)

func double(primeNum float64, ch chan<- *PrimeNumber) {
	fmt.Println("Starting mapper:", primeNum)
	time.Sleep(2 * time.Second)
	ch <- &PrimeNumber{Num: primeNum * 2}

	// some other code
}

func reducer(ch <-chan *PrimeNumber) <-chan float64 {
	resultCh := make(chan float64)

	go func() {
		result := float64(0)
		for val := range ch {
			result += val.Num
		}
		resultCh <- result
	}()

	return resultCh
}

type PrimeNumber struct { // >64 KiB
	Num  float64           // 8 bytes
	data [100_000_000]byte // <64 KiB
}

func main() {
	// Map-Reduce Pipeline => primes => double primes (map) => sum of doubled primes (reduce)

	primes := []float64{1, 2, 3, 5, 7}

	doubledCh := make(chan *PrimeNumber) // Buffered vs unbuffered channel

	var wg sync.WaitGroup
	for _, number := range primes {
		wg.Add(1)
		go func(el float64) {
			defer wg.Done()
			double(el, doubledCh)
		}(number)
	}

	go func() {
		defer close(doubledCh)
		wg.Wait()
	}()

	resultCh := reducer(doubledCh)

	fmt.Println("Result:", <-resultCh)
}
