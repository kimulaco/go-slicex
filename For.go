package slicex

func (s *Slicex[T]) For (fn func (v T, i int)) {
	for i, v := range s.Data {
		fn(v, i)
	}
}

func (s *Slicex[T]) Filter (fn func (v T, i int) bool) {
	d := make([]T, 0, s.Length())

	for i, v := range s.Data {
		if fn(v, i) {
			d = append(d, v)
		}
	}

	s.Data = d
}

func (s *Slicex[T]) Map (fn func (v T, i int) T) {
	d := make([]T, 0, s.Length())

	for i, v := range s.Data {
		d = append(d, fn(v, i))
	}

	s.Data = d
}
