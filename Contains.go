package slicex

func (s *Slicex[T]) Contains (v T) bool {
	return s.IndexOf(v) != -1
}
