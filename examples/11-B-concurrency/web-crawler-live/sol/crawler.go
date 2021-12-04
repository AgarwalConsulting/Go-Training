// Example from: https://tour.golang.org/concurrency/10
package main

import (
	"fmt"
	"sync"
)

type Fetcher interface { // Web Scraper abstraction
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Maps are NOT safe for concurrent safe
// var visitedURL map[string]bool // Default Value: <nil>

// Initialize a map?

// var visitedURL = map[string]bool{} // Map-literal syntax
var visitedURL = make(map[string]bool) // Make

var mut sync.Mutex

func isVisitedAndSet(url string) bool {
	mut.Lock()
	defer mut.Unlock()
	val := visitedURL[url]
	visitedURL[url] = true
	return val
}

// var results []Result // Default Value: <nil>

// Crawl uses fetcher to recursively crawl (Depth-first recursive crawl)
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan<- Result) { // Send-only channel
	// TODO 1: Fetch URLs in parallel. - Done!
	// TODO 2: Don't fetch the same URL twice. - Done!
	// TODO 3: Collect all the results in a []Result
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	if isVisitedAndSet(url) { // Atomic
		return
	}

	body, urls, err := fetcher.Fetch(url)
	res := Result{URL: url, body: body, IsFound: err != nil}
	ch <- res // Sending to the channel; blocks until there is a receiver
	// results = append(results, res)

	if err != nil {
		// fmt.Println(err)
		return
	}
	// fmt.Printf("found: %s %q\n", url, body)

	var wg sync.WaitGroup // Counter: 0

	for _, u := range urls {
		wg.Add(1) // Increment the counter by 1

		go func(url string) {
			defer wg.Done()                  // Decrement the counter by 1
			Crawl(url, depth-1, fetcher, ch) // Captures
		}(u)
	}

	wg.Wait() // Blocks until the counter reaches 0 again

	return
}

func main() { // Main goroutine
	// var ch chan Result // Default: <nil>
	ch := make(chan Result)
	var results []Result

	// go func() {
	// }()

	go func() {
		defer close(ch)
		defer fmt.Println("Closing channel...")
		defer fmt.Println("---------Done Crawling---------")

		Crawl("https://golang.org/", 10, fetcher, ch) // Producer
	}()

	for r := range ch { // Blocks until next value is sent over the channel; terminates the loop if channel is closed!
		results = append(results, r)
	}

	fmt.Println(results)

	// val, ok := <-ch // Receiving from a closed channel doesn't block and returns the default value
	// if ok {
	// 	fmt.Println(val) // {false "" ""}
	// } else {
	// 	fmt.Println("Channel is closed!")
	// }

	// ch <- val // Sending to a closed channel => panic!
	// close(ch) // Closing a closed channel => panic!
}
