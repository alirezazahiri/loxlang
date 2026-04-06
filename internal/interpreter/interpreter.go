package interpreter

import (
	"loxlang/internal/environment"
)

type Interpreter struct {
	env *environment.Environment
}

func New() *Interpreter {
	return &Interpreter{
		env: environment.New(),
	}
}
