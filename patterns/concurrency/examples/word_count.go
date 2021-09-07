package main

import (
	"fmt"
	"strings"
	"sync"
)

type Emit struct {
	word  string
	count int
}

type FreqMap map[string]int

func mapper(id int, sentences <-chan string, ch chan<- Emit, wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Printf("Mapper %d is terminating...\n", id)
	fmt.Printf("Mapper %d is starting...\n", id)

	for sentence := range sentences {
		fm := FreqMap{}
		words := strings.Split(sentence, " ")

		for _, word := range words {
			word = strings.TrimSpace(word)
			if len(word) > 0 {
				fm[word] += 1
			}
		}

		for word, count := range fm {
			emitValue := Emit{word, count}
			fmt.Println("\tEmiting: ", emitValue, "from: Mapper", id)
			ch <- emitValue
		}
	}
}

func reducer(ch <-chan Emit) <-chan FreqMap {
	result := make(chan FreqMap)

	go func() {
		fm := FreqMap{}

		for emit := range ch {
			fmt.Println("\tConsuming:", emit)
			fm[emit.word] += emit.count
		}

		result <- fm
	}()

	return result
}

func splitIntoSentences(text string) <-chan string {
	result := make(chan string)

	go func() {
		var sentence = []rune{}

		for _, el := range text {
			if el == '.' || el == '?' || el == '!' || el == ';' {
				if len(sentence) > 0 {
					result <- string(sentence)
				}
				sentence = nil
			} else {
				sentence = append(sentence, el)
			}
		}

		if len(sentence) > 0 {
			result <- string(sentence)
		}
		close(result)
	}()

	return result
}

func main() {
	text := `Parallelism is about doing things in parallel that can also be done sequentially. A common example is counting the frequency of letters. Create a function that returns the total frequency of each letter in a list of texts and that employs parallelism.

	Go supports concurrency via goroutines which are started with the go keyword. It is a simple, lightweight and elegant way to provide concurrency support and is one of the greatest strengths of the language.

	You may notice that while this exercise is called Parallel letter frequency you don't see the term Parallel used very often in Go. Gophers prefer to use the term Concurrent to describe the management of multiple independent goroutines. Although these terms are often used interchangeably Gophers like to be technically correct and use concurrent when discussing the seemingly simultaneous executions of goroutines.

	While we can plan for our programs to run in parallel, and at times they may appear to run in parallel, without strict knowledge of the execution context of our code all we can guarantee is that processes will run concurrently. In other words they may be executing sequentially faster than we can distinguish but not strictly simultaneously.`

	sentenceCh := splitIntoSentences(text)

	emitCh := make(chan Emit) // Unbuffered
	// emitCh := make(chan Emit, 100) // Buffered

	var wg sync.WaitGroup
	wg.Add(3)

	go mapper(1, sentenceCh, emitCh, &wg)
	go mapper(2, sentenceCh, emitCh, &wg)
	go mapper(3, sentenceCh, emitCh, &wg)

	freqMapFuture := reducer(emitCh)

	wg.Wait()

	close(emitCh)

	fmt.Printf("%#v\n", <-freqMapFuture)
}
