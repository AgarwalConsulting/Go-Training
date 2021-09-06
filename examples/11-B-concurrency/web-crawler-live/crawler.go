// Example from: https://tour.golang.org/concurrency/10
package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// var visitedURL map[string]bool

// var visitedURL = map[string]bool{} // Map-literal syntax
// var visitedURL = make(map[string]bool)
// var mut sync.Mutex

// var vm VisitedMap // Default: {m: nil}
var vm = NewVisitedMap()

// var results []Result // Default: nil

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
// TODO 1: Fetch URLs in parallel. - Done
func Crawl(url string, depth int, fetcher Fetcher, ch chan<- Result) { // Send-only
	// TODO 2: Don't fetch the same URL twice. - Done
	// TODO 3: Collect all the results in a []Result
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	if vm.IsVisitedAndSet(url) { // Atomic
		return
	}

	body, urls, err := fetcher.Fetch(url)

	res := Result{URL: url, body: body, IsFound: err == nil}
	// results = append(results, res)
	ch <- res // Sending

	if err != nil {
		// fmt.Println(err)
		return
	}
	// fmt.Printf("found: %s %q\n", url, body)

	var wg sync.WaitGroup // Initial: 0

	for _, u := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			Crawl(url, depth-1, fetcher, ch)
		}(u)
	}
	wg.Wait()
}

func main() {
	// var ch chan Result
	ch := make(chan Result)

	go func() {
		for res := range ch { // blocks until the next value; ranges stops if my channel is closed
			// res := <-ch // Receiving
			fmt.Printf("%#v\n", res)
		}
	}()

	// go func() {
	Crawl("https://golang.org/", 10, fetcher, ch)
	// 	close(ch)
	// 	fmt.Println("Completed crawling!")
	// }()

	// Receiving from a closed channel
	// v, ok := <-ch // Doesn't block
	// if ok {
	// 	fmt.Printf("%T %#v\n", v, v) // ?
	// } else {
	// 	fmt.Println("Channel is closed!")
	// }

	// close(ch)
	// ch <- v // Sending to a closed channel
}
