package scope

type Scope map[string]bool

func (s Scope) Put(key string, value bool) {
	s[key] = value
}
func (s Scope) Get(key string) (bool, bool) {
	v, ok := s[key]
	return v, ok
}
func (s Scope) ContainsKey(key string) bool {
	_, ok := s[key]
	return ok
}
