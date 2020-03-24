package functions

import (
	"fmt"
	"os"
	"testing"
)

// go test								Run tests for current package.
// go test ./... 						Tests current and all subpackages.
// go test --bench=. --benchmem			Run benchmarks including memory.
// go test -coverprofile=coverage.out	Output coverage profile in a file.
// go tool cover -func=coverage.out		Function wise coverage from the output file.
// go tool cover -html=coverage.out		Detailed coverage in Browser

func SnT(m *testing.M) int {
	defer fmt.Println("Teardown...")

	fmt.Println("Setup...")

	return m.Run()
}

func TestMain(m *testing.M) {
	os.Exit(SnT(m))
}

func TestSimpleFac(t *testing.T) {
	input := 5
	output := 120

	r := fac(input)

	if output != r {
		t.Fatalf("Factorial %v: Expected %v, Got %v ", input, output, r)
	}
}

func TestFac(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{0, 1},
		{1, 1},
		{5, 120},
	}

	for _, tc := range tests {
		r := fac(tc.input)
		if tc.output != r {
			t.Fatalf("Factorial %v: Expected %v, Got %v ", tc.input, tc.output, r)
		}
	}
}

func TestFacr(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{0, 1},
		{1, 1},
		{5, 120},
	}

	for _, tc := range tests {
		r := facr(tc.input)
		if tc.output != r {
			t.Fatalf("Factorial %v: Expected %v, Got %v ", tc.input, tc.output, r)
		}
	}
}

var result int

func BenchmarkFac20(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		fac(20)
	}
}

// https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
func BenchmarkFacr20(b *testing.B) {
	var innerResult int
	for i := 0; i <= b.N; i++ {
		innerResult = facr(20)
	}
	result = innerResult
}
