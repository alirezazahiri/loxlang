package environment

func (e *Environment) Define(name string, value any) {
	e.values[name] = value
}

func (e *Environment) Get(name string) (any, bool) {
	value, ok := e.values[name]
	return value, ok
}
