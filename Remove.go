package slicex

func RemoveIndex[T any](vv []T, i int) []T {
	if i < 0 || i >= len(vv) {
		return vv
	}

	return append(vv[:i], vv[i+1:]...)
}

func RemoveValue[T any](vv []T, v T) []T {
	i := IndexOf(vv, v)

	if i == -1 {
		return vv
	}

	return RemoveIndex(vv, i)
}
