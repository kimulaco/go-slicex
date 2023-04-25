package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	s1 := New([]int{1, 2, 3})
	assert.Equal(t, s1.Contains(2), true)

	s2 := New([]int{1, 2, 3})
	assert.Equal(t, s2.Contains(4), false)
}
