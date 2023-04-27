package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveIndex(t *testing.T) {
	assert.Equal(t, []int{1, 3}, RemoveIndex([]int{1, 2, 3}, 1))

	assert.Equal(t, []int{1, 2, 3}, RemoveIndex([]int{1, 2, 3}, 4))
	assert.Equal(t, []int{1, 2, 3}, RemoveIndex([]int{1, 2, 3}, -1))
}

func TestRemoveValue(t *testing.T) {
	assert.Equal(t, []int{1, 3}, RemoveValue([]int{1, 2, 3}, 2))

	assert.Equal(t, []int{1, 2, 3}, RemoveValue([]int{1, 2, 3}, 4))
}
