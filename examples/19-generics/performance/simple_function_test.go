package main

import "testing"

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(10, 12)
	}
}

func BenchmarkPrintMeG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addG(10, 12)
	}
}
