package main

import (
	"fmt"
	"strings"
	"sync"
)

type WordCountMap map[string]int

type Emit struct {
	Word  string
	Count int
}

func splitSentence(text string) []string {
	output := []string{}

	sentence := []rune{}
	for _, el := range text {
		if strings.Contains(`.,?!;()"`, string(el)) {
			if len(sentence) > 0 {
				output = append(output, string(sentence))
			}
			sentence = nil
		} else {
			sentence = append(sentence, el)
		}
	}

	if len(sentence) > 0 {
		output = append(output, string(sentence))
	}

	return output
}

func mapper(sentence string) <-chan Emit {
	output := make(chan Emit)

	go func() {
		words := strings.Split(sentence, " ")
		wcm := WordCountMap{}
		// wcm := make(WordCountMap)

		for _, word := range words {
			wcm[word]++
		}

		for word, count := range wcm {
			output <- Emit{word, count}
		}
		close(output)
	}()

	return output
}

func fanIn[T any](chs []<-chan T) <-chan T {
	output := make(chan T)

	var wg sync.WaitGroup

	for _, ch := range chs {
		wg.Add(1)
		go func(ch <-chan T) {
			defer wg.Done()
			for el := range ch {
				output <- el
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

func reducer(chs []<-chan Emit) WordCountMap {
	emitCh := fanIn(chs)

	output := WordCountMap{}

	for emit := range emitCh {
		output[emit.Word] += emit.Count
	}

	return output
}

func main() {
	text := `Package sync provides basic synchronization primitives such as mutual exclusion locks. Other than the Once and WaitGroup types, most are intended for use by low-level library routines. Higher-level synchronization is better done via channels and communication.

	Values containing the types defined in this package should not be copied.`

	// Chunking: Split into sentences
	sentences := splitSentence(text) // Type: []string

	var wordCountChs []<-chan Emit
	// Mapper: Word count from each sentence
	for _, sentence := range sentences {
		wordCountCh := mapper(sentence) // Type: <-chan Emit
		wordCountChs = append(wordCountChs, wordCountCh)
	}

	// Reducer: Add up all the word count
	wordCountMap := reducer(wordCountChs)

	fmt.Printf("%#v\n", wordCountMap)
}
