package slicex

type Slicex[T any] struct {
	Data []T
}

func New[T any](d []T) *Slicex[T] {
	return &Slicex[T]{
		Data: d,
	}
}
