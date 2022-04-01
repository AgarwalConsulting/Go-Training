package calculator_test

import (
	"testing"

	"algogrit.com/calculator"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	expected := 5
	actual := calculator.Add(2, 3)

	assert.Equal(t, expected, actual)
	// if expected != actual {
	// 	t.Fail()
	// }
}

func TestMul(t *testing.T) {
	testCase := struct {
		input    [2]int
		expected int
	}{
		[2]int{2, 3}, 6,
	}

	actual := calculator.Mul(testCase.input[0], testCase.input[1])
	assert.Equal(t, testCase.expected, actual)
	// if expected != actual {
	// 	t.Log("Expected: ", expected)
	// 	t.Log("Actual: ", actual)
	// 	t.Fail()
	// }
}

func TestMulCases(t *testing.T) {
	testCases := []struct {
		input    [2]int
		expected int
	}{
		{[2]int{2, 5}, 10},
		{[2]int{10, 5}, 50},
		{[2]int{5, 5}, 25},
		{[2]int{12, 6}, 72},
	}

	for _, testCase := range testCases {
		actual := calculator.Mul(testCase.input[0], testCase.input[1])
		if testCase.expected != actual {
			t.Log("Expected: ", testCase.expected)
			t.Log("Actual: ", actual)
			t.Fatal()
		}
	}
}
