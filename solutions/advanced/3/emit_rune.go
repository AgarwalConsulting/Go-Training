package letter

import (
	"fmt"
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

func (f FreqMap) String() string {
	var output string = "{ "

	for char, count := range f {
		output += fmt.Sprintf("%v = %c : %d | ", char, char, count)
	}

	output += " }"

	return output
}

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func sendRunes(str string, ch chan<- rune, wg *sync.WaitGroup) {
	for _, char := range str {
		ch <- char
	}
	wg.Done()
}

func closeChannel(wg *sync.WaitGroup, ch chan<- rune) {
	wg.Wait()
	close(ch)
}

// ConcurrentFrequency counts the frequency of each rune in a given slice of text and returns this
// data as a FreqMap.
func ConcurrentFrequency(s []string) FreqMap {
	var wg sync.WaitGroup
	freq := make(FreqMap)
	ch := make(chan rune, 10)
	for _, str := range s {
		wg.Add(1)
		go sendRunes(str, ch, &wg)
	}

	go closeChannel(&wg, ch)

	for char := range ch {
		freq[char]++
	}

	return freq
}
