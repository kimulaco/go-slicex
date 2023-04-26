package slicex

import (
	"reflect"
)

func IndexOf[T any] (vv []T, v T) int {
	for i, d := range vv {
		if reflect.DeepEqual(d, v) {
			return i
		}
	}

	return -1
}
