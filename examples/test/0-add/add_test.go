package add

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddFn(t *testing.T) {
	i := Add(2, 3)

	// if i != 5 {
	// 	fmt.Println(i, " isn't equal to ", 5)
	// 	t.Fail()
	// }

	assert.Equal(t, 5, i)
}
