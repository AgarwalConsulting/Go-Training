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

func calculateFreq(cfm *ConcurrentFreqMap, str string) {
	defer cfm.wg.Done()
	for _, char := range str {
		cfm.add(char)
	}
}

type ConcurrentFreqMap struct {
	freqMap FreqMap
	mut     sync.Mutex
	wg      sync.WaitGroup
}

func NewConcurrentFreqMap() *ConcurrentFreqMap {
	return &ConcurrentFreqMap{freqMap: make(map[rune]int)}
}

func (cfm *ConcurrentFreqMap) add(r rune) {
	cfm.mut.Lock()
	defer cfm.mut.Unlock()
	cfm.freqMap[r] = cfm.freqMap[r] + 1
}

// ConcurrentFrequency counts the frequency of each rune in a given slice of text and returns this
// data as a FreqMap.
func ConcurrentFrequency(s []string) FreqMap {
	cfm := NewConcurrentFreqMap()
	for _, str := range s {
		cfm.wg.Add(1)
		go calculateFreq(cfm, str)
	}
	cfm.wg.Wait()
	return cfm.freqMap
}
