package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Resp struct {
	Url   string
	Body  []byte
	Error error
}

func fetch(url string) <-chan Resp {
	respChan := make(chan Resp)
	go func() {
		resp, err := http.Get(url)
		if err != nil {
			respChan <- Resp{Url: url, Body: nil, Error: err}
		} else {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			respChan <- Resp{Url: url, Body: body, Error: err}
		}
	}()
	return respChan
}

func respPrinter(startTime time.Time, resp Resp) {
	duration := time.Now().Sub(startTime)
	if resp.Error != nil {
		fmt.Printf("error: %#v", resp.Error)
		return
	}
	fmt.Printf(
		"response from %s length: %d bytes in %f seconds\n",
		resp.Url,
		len(resp.Body),
		duration.Seconds())
}

func reducer(futures []<-chan Resp) chan Resp {
	aggregate := make(chan Resp)
	var wg sync.WaitGroup

	wg.Add(len(futures))
	for _, f := range futures {
		future := f
		go func() {
			resp := <-future
			aggregate <- resp
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(aggregate)
	}()

	return aggregate
}

func main() {
	startTime := time.Now()

	urls := []string{
		"https://tarkalabs.com/",
		"https://algogrit.com/",
		"https://google.com/",
	}

	var futures []<-chan Resp

	for _, url := range urls {
		futures = append(futures, fetch(url))
	}

	fmt.Println("doing other busy things here.")

	aggregate := reducer(futures)

	timeOut := time.After(1 * time.Second)

	for {
		select {
		case resp, ok := <-aggregate:
			if ok {
				respPrinter(startTime, resp)
			}
		case <-timeOut:
			fmt.Println("Timeout!")
			return
		}
	}
}
