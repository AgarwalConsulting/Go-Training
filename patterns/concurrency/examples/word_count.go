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

func mapper(sentences <-chan string, ch chan<- Emit, wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("Mapper is terminating...")
	fmt.Println("Mapper is starting...")

	for sentence := range sentences {
		fm := FreqMap{}
		words := strings.Split(sentence, " ")

		for _, word := range words {
			fm[word] += 1
		}

		for word, count := range fm {
			if len(word) > 0 {
				ch <- Emit{word, count}
			}
		}
	}
}

func reducer(ch <-chan Emit) <-chan FreqMap {
	result := make(chan FreqMap)

	go func() {
		fm := FreqMap{}

		for emit := range ch {
			fmt.Println("\tEmit:", emit)
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
	text := "what are you doing? Do you like ice-cream? I like ice-cream."

	sentenceCh := splitIntoSentences(text)

	emitCh := make(chan Emit) // Unbuffered

	var wg sync.WaitGroup
	wg.Add(3)

	go mapper(sentenceCh, emitCh, &wg)
	go mapper(sentenceCh, emitCh, &wg)
	go mapper(sentenceCh, emitCh, &wg)

	freqMapFuture := reducer(emitCh)

	wg.Wait()

	close(emitCh)

	fmt.Printf("%#v\n", <-freqMapFuture)
}
