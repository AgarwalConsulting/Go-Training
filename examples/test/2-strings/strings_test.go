package test

import (
	"strings"
	"testing"
)

var split []string

func BenchmarkSplit(b *testing.B) {
	var localSplit []string
	for i := 0; i < b.N; i++ {
		localSplit = strings.Split("one,two,three,four,five,six,seven", ",")
	}

	split = localSplit
}
