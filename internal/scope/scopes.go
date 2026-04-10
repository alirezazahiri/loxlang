package scope

type Scopes []Scope

func (s Scopes) IsEmpty() bool {
	return len(s) == 0
}
func (s Scopes) Peek() Scope {
	if len(s) == 0 {
		return nil
	}
	return s[len(s)-1]
}
func (s Scopes) Pop() Scope {
	last := s.Peek()
	s = s[:len(s)-1]
	return last
}
func (s Scopes) Size() int {
	return len(s)
}
func (s Scopes) Get(i int) Scope {
	return s[i]
}
