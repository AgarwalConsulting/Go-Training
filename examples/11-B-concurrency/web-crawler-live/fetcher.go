package main

import (
	"fmt"
	"math/rand"
	"time"
)

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func init() {
	rand.Seed(time.Now().Unix())
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	fakeLatency := time.Duration(rand.Intn(700)) * time.Millisecond

	time.Sleep(fakeLatency)

	fmt.Println("--Latency:", fakeLatency, "while fetching", url)

	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}
