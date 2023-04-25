package slicex

import (
	"reflect"
)

func (s *Slicex[T]) IndexOf (v T) int {
	for i, d := range s.Data {
		if reflect.DeepEqual(d, v) {
			return i
		}
	}

	return -1
}
