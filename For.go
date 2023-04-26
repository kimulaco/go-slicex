package slicex

func  ForEach[T any] (vv []T, fn func (v T, i int)) {
	for i, v := range vv {
		fn(v, i)
	}
}

func Filter[T any] (vv []T, fn func (v T, i int) bool) []T {
	d := make([]T, 0, len(vv))

	for i, v := range vv {
		if fn(v, i) {
			d = append(d, v)
		}
	}

	return d
}

func Map[T any] (vv []T, fn func (v T, i int) T) []T {
	d := make([]T, 0, len(vv))

	for i, v := range vv {
		d = append(d, fn(v, i))
	}

	return d
}

func Find[T any] (vv []T, fn func (v T, i int) bool) (T, bool) {
	for i, v := range vv {
		if fn(v, i) {
			return v, true
		}
	}

	var notFound T
	return notFound, false
}
