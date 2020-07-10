package main_test

import (
	"testing"

	example "github.com/AgarwalConsulting/Go-Training"
)

func TestAdd(t *testing.T) {
	result := example.Add(12, 14)

	if result != 26 {
		t.Fail()
	}
}

func BenchmarkAddTwoPlusTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		example.Add(2, 2)
	}
}
