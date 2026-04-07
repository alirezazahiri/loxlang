package environment

func (e *Environment) Define(name string, value any) {
	e.values[name] = value
}

func (e *Environment) Assign(name string, value any) bool {
	_, ok := e.values[name]
	if ok {
		e.values[name] = value
		return true
	}

	if e.enclosing != nil {
		return e.enclosing.Assign(name, value)
	}

	return false
}

func (e *Environment) Get(name string) (any, bool) {
	value, ok := e.values[name]

	if ok {
		return value, true
	}

	if e.enclosing != nil {
		return e.enclosing.Get(name)
	}

	return nil, false
}
