package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	assert.Equal(t, Contains([]int{1, 2, 3}, 2), true)
	assert.Equal(t, Contains([]int{1, 2, 3}, 4), false)
}
