package slicex

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	d := []int{1, 2, 3}
	s := New(d)

	if !reflect.DeepEqual(s.Data, d) {
		t.Errorf("New(%v) = %v, want %v", d, s.Data, d)
	}
}
