package slicex

func (s *Slicex[T]) Add (v T) {
	s.Data = append(s.Data, v)
}
