package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveValue(t *testing.T) {
	s1 := New([]int{1, 2, 3})
	s1.RemoveValue(2)
	assert.Equal(t, []int{1, 3}, s1.Data)
}

func TestRemoveIndex(t *testing.T) {
	s1 := New([]int{1, 2, 3})
	s1.RemoveIndex(1)
	assert.Equal(t, []int{1, 3}, s1.Data)
}
