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

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, m ConcurrentMap, ch chan<- Result) {
	// TODO: Fetch URLs in parallel. - Done
	// TODO: Don't fetch the same URL twice.
	if depth <= 0 {
		return
	}

	if m.CheckAndSet(url) {
		// fmt.Println("\tskipping: ", url)
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		// fmt.Println(err)
		ch <- Result{URL: url}
		return
	}
	// fmt.Printf("found: %s %q\n", url, body)
	ch <- Result{true, url, body}

	var wg sync.WaitGroup

	for _, u := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			Crawl(u, depth-1, fetcher, m, ch)
		}(u)
	}
	wg.Wait()
	return
}

func main() {
	var m = NewConcurrentMap()
	ch := make(chan Result)

	go Crawl("https://golang.org/", 10, fetcher, m, ch)

	// var results []Result

	for res := range ch { // Blocks until it receives next value;
		fmt.Printf("%T %#v\n", res, res) // Result
	}
}
