package calculator_test

import (
	"testing"

	calculator "github.com/Chennai-Golang/101-workshop/examples/test/0-calculator"
)

func TestAdd(t *testing.T) {
	expected := 5
	actual := calculator.Add(3, 2)
	if expected != actual {
		t.Log("Expected: ", expected)
		t.Log("Actual: ", actual)
		t.Fail()
	}
}

func TestMul(t *testing.T) {
	testCases := []struct {
		input    [2]int
		expected int
	}{
		{[2]int{10, 5}, 50},
		{[2]int{6, 3}, 18},
		{[2]int{25, 100}, 2500},
	}

	for _, testCase := range testCases {
		actual := calculator.Mul(testCase.input[0], testCase.input[1])
		if testCase.expected != actual {
			t.Log("Expected: ", testCase.expected)
			t.Log("Actual: ", actual)
			t.Fatal("Failed!")
		}
	}
}
