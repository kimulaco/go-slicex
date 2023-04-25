package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFor(t *testing.T) {
	s1 := New([]int{0, 1, 4, 9, 16, 25, 36})
	c := 0
	s1.For(func (v int, i int) {
		assert.Equal(t, i, c)
		assert.Equal(t, v, c * c)
		c++
	})
}

func TestFilter(t *testing.T) {
	s1 := New([]int{0, 1, 2, -3, -4, 5, 7})
	s1.Filter(func (v int, i int) bool {
		return v == i
	})
	assert.Equal(t, s1.Data, []int{0, 1, 2, 5})
}

func TestMap(t *testing.T) {
	s1 := New([]int{0, 10, 20, 30, 40, 50, 60})
	s1.Map(func (v int, i int) int {
		return v * i
	})
	assert.Equal(t, s1.Data, []int{0, 10, 40, 90, 160, 250, 360})
}
