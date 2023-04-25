package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	s1 := New([]int{1, 2, 3})
	s1.Add(4)
	assert.Equal(t, s1.Data, []int{1, 2, 3, 4})
}
