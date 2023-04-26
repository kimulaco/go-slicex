package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForEach(t *testing.T) {
	s1 := []int{0, 1, 4, 9, 16, 25, 36}
	c := 0
	ForEach(s1, func (v int, i int) {
		assert.Equal(t, i, c)
		assert.Equal(t, v, c * c)
		c++
	})
}

func TestFilter(t *testing.T) {
	s1 := []int{0, 1, 2, -3, -4, 5, 7}
	s1 = Filter(s1, func (v int, i int) bool {
		return v == i
	})
	assert.Equal(t, s1, []int{0, 1, 2, 5})
}

func TestMap(t *testing.T) {
	s1 := []int{0, 10, 20, 30, 40, 50, 60}
	s1 = Map(s1, func (v int, i int) int {
		return v * i
	})
	assert.Equal(t, s1, []int{0, 10, 40, 90, 160, 250, 360})
}
