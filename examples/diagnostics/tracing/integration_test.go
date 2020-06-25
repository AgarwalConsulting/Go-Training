package trace

// Credits: https://medium.com/@cep21/using-go-1-10-new-trace-features-to-debug-an-integration-test-1dc39e4e812d
// Code: https://github.com/cep21/trace-110/blob/master/integration_test.go

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"
)

func takeCPU(start time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	j := 3
	for time.Since(start) < time.Second {
		for i := 1; i < 1000000; i++ {
			j *= i
		}
	}
	fmt.Println(j)
}

func takeTimeOnly(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 3)
}

func takeIO(start time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	errCount := 0
	for time.Since(start) < time.Second*4 {
		_, err := http.Get("https://www.google.com")
		if err != nil {
			errCount++
		}
	}
	fmt.Println(errCount)
}

func startServer() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	start := time.Now()
	go takeCPU(start, &wg)
	go takeTimeOnly(&wg)
	go takeIO(start, &wg)
	wg.Wait()
}

func TestServer(t *testing.T) {
	startServer()
}

func BenchmarkServer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		startServer()
	}
}
