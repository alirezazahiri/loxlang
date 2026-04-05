package interpreter

import (
	"loxlang/pkg/loxerror"
)

type Interpreter struct {
	error *loxerror.LoxError
}

func New() *Interpreter {
	return &Interpreter{
		error: nil,
	}
}
