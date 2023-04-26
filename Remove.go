package slicex

func RemoveIndex[T any] (vv []T, i int) []T {
	return append(vv[:i], vv[i+1:]...)
}

func RemoveValue[T any] (vv []T, v T) []T {
	i := IndexOf(vv, v)
	return RemoveIndex(vv, i)
}
