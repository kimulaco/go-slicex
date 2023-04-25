package slicex

import (
	"testing"
)

func TestContains(t *testing.T) {
	d1 := []int{1, 2, 3}
	s1 := New(d1)
	c1 := s1.Contains(2)
	if c1 != true {
		t.Errorf("New(%v) = %t, want %t", d1, c1, true)
	}

	d2 := []int{1, 2, 3}
	s2 := New(d2)
	i2 := s2.Contains(4)
	if i2 != false {
		t.Errorf("New(%v) = %t, want %t", d2, i2, false)
	}
}
