package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLength(t *testing.T) {
	s1 := New([]int{1, 2, 3})
	assert.Equal(t, s1.Length(), 3)

	s2 := New([]int{})
	assert.Equal(t, s2.Length(), 0)
}
