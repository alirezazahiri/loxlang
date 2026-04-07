package environment

type Environment struct {
	enclosing *Environment
	values    map[string]any
}

func New(enclosing *Environment) *Environment {
	return &Environment{
		values:    map[string]any{},
		enclosing: enclosing,
	}
}