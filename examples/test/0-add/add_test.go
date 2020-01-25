package add

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddFn(t *testing.T) {
	i := Add(2, 3)

	assert.Equal(t, 5, i)
}
