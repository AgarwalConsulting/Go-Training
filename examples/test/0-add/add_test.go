package add

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	expected := 5
	output := Add(2, 3)

	// if output != expected {
	// 	t.Log("Expected: ", expected, " Got: ", output)
	// 	t.Fail()
	// }

	assert.Equal(t, expected, output)
}
