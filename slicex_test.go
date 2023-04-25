package slicex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	d1 := []int{1, 2, 3}
	s1 := New(d1)
	assert.Equal(t, s1.Data, d1)
}
