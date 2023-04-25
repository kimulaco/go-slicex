package slicex

func (s *Slicex[T]) Length () int {
	return len(s.Data)
}
