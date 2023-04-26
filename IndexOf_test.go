package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexOf(t *testing.T) {
	assert.Equal(t, IndexOf([]int{1, 2, 3}, 2), 1)
	assert.Equal(t, IndexOf([]int{1, 2, 3}, 4), -1)
}
