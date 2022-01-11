package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func mapper(num int) <-chan float64 {
	ch := make(chan float64)

	go func() {
		defer fmt.Println("Mapped:", num)
		dur := time.Duration(10-num) * time.Second

		fmt.Println("Mapping...", num, dur)
		time.Sleep(dur) // Simulating latency on my mapper
		ch <- math.Sqrt(float64(num))
	}()

	return ch
}

func fanIn(inChs []<-chan float64) <-chan float64 {
	resCh := make(chan float64)

	go func() {
		var wg sync.WaitGroup

		for _, ch := range inChs {
			wg.Add(1)
			go func(mapCh <-chan float64) {
				defer wg.Done()
				mappedData := <-mapCh
				resCh <- mappedData
			}(ch)
		}

		wg.Wait()
		close(resCh)
	}()

	return resCh
}

func reduce(inChs []<-chan float64) <-chan float64 {
	resultCh := make(chan float64)

	go func() {
		defer fmt.Println("Terminating reducer...")
		var sum float64
		// for idx, mapperCh := range inChs {
		// 	data := <-mapperCh
		// 	fmt.Println("Reducing...", idx, data)
		// 	sum += data
		// }

		fannedInCh := fanIn(inChs)
		for data := range fannedInCh {
			sum += data
		}

		resultCh <- sum
	}()

	return resultCh
}

func main() {
	fmt.Println("Starting map-reduce...")

	primes := []int{1, 2, 3, 5, 7}

	var mapperChs []<-chan float64

	for _, el := range primes {
		mapperChs = append(mapperChs, mapper(el))
	}

	resultCh := reduce(mapperChs)

	fmt.Println("Result:", <-resultCh)
}
