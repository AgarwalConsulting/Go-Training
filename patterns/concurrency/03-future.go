package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Resp struct {
	Body  []byte
	Error error
}

func fetch(url string) <-chan Resp {
	respChan := make(chan Resp)
	go func() {
		resp, err := http.Get(url)
		if err != nil {
			respChan <- Resp{Body: nil, Error: err}
		} else {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			respChan <- Resp{Body: body, Error: err}
		}
		close(respChan)
	}()
	return respChan
}

func main() {
	startTime := time.Now()
	future := fetch("https://algogrit.com/") // async fetch

	fmt.Println("doing other busy things here.")

	body := <-future // body := await future
	duration := time.Now().Sub(startTime)
	if body.Error != nil {
		fmt.Printf("error: %#v", body.Error)
		return
	}
	fmt.Printf(
		"response length: %d bytes in %f seconds\n",
		len(body.Body),
		duration.Seconds())
}
