package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexOf(t *testing.T) {
	s1 := New([]int{1, 2, 3})
	assert.Equal(t, s1.IndexOf(2), 1)

	s2 := New([]int{1, 2, 3})
	assert.Equal(t, s2.IndexOf(4), -1)
}
