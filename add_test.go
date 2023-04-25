package slicex

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	d1 := []int{1, 2, 3}
	_d1 := []int{1, 2, 3, 4}
	s1 := New(d1)
	s1.Add(4)
	if !reflect.DeepEqual(s1.Data, _d1) {
		t.Errorf("New(%v) = %v, want %v", _d1, s1.Data, _d1)
	}
}
