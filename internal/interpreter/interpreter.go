package interpreter

import (
	"loxlang/internal/environment"
)

type Interpreter struct {
	env *environment.Environment
}

func New(globals *environment.Environment) *Interpreter {
	return &Interpreter{
		env: environment.New(globals),
	}
}
