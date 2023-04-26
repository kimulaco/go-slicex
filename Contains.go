package slicex

func Contains[T any] (vv []T, v T) bool {
	return IndexOf(vv, v) != -1
}
