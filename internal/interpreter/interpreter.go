package interpreter

import (
	"loxlang/internal/environment"
)

type Interpreter struct {
	env *environment.Environment
}

func New(env *environment.Environment) *Interpreter {
	return &Interpreter{
		env: env,
	}
}
