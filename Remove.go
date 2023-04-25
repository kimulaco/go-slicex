package slicex

func (s *Slicex[T]) RemoveValue (v T) {
	i := s.IndexOf(v)
	s.RemoveIndex(i)
}

func (s *Slicex[T]) RemoveIndex (i int) {
	s.Data = append(s.Data[:i], s.Data[i+1:]...)
}
