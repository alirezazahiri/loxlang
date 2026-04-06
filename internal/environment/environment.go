package environment

type Environment struct {
	values map[string]any
}

func New() *Environment {
	return &Environment{
		values: map[string]any{},
	}
}