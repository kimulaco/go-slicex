package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveIndex(t *testing.T) {
	assert.Equal(t, RemoveIndex([]int{1, 2, 3}, 1), []int{1, 3})
}

func TestRemoveValue(t *testing.T) {
	assert.Equal(t, RemoveValue([]int{1, 2, 3}, 2), []int{1, 3})
}
