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

// ConcurrentFrequency counts the frequency of each rune in a given slice of text and returns this
// data as a FreqMap.
func ConcurrentFrequencyWO(s []string) FreqMap {
	fm := FreqMap{}
	var wg sync.WaitGroup
	var mut sync.Mutex

	for _, str := range s {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			for _, char := range s {
				mut.Lock()
				fm[char]++
				mut.Unlock()
			}
		}(str)
	}
	wg.Wait()

	return fm
}

func ConcurrentFrequencyWC(s []string) FreqMap {
	freqMap := FreqMap{}
	ch := make(chan FreqMap, len(s))
	for idx := range s {
		go func(str string) {
			ch <- Frequency(str)
		}(s[idx])
	}

	for i := 0; i < len(s); i++ {
		fm := <-ch
		for key, value := range fm {
			freqMap[key] += value
		}
	}
	close(ch)
	return freqMap
}
