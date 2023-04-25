package slicex

import (
	"testing"
)

func TestIndexOf(t *testing.T) {
	d1 := []int{1, 2, 3}
	s1 := New(d1)
	i1 := s1.IndexOf(2)
	if i1 != 1 {
		t.Errorf("New(%v) = %c, want %d", d1, i1, 1)
	}

	d2 := []int{1, 2, 3}
	s2 := New(d2)
	i2 := s2.IndexOf(4)
	if i2 != -1 {
		t.Errorf("New(%v) = %c, want %d", d2, i2, -1)
	}
}
