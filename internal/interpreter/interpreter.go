package interpreter

import (
	"loxlang/internal/ast"
	"loxlang/internal/environment"
)

type Interpreter struct {
	env     *environment.Environment
	globals *environment.Environment
	locals  map[ast.Expr]int
}

func New(globals *environment.Environment) *Interpreter {
	return &Interpreter{
		env:     environment.New(globals),
		globals: globals,
		locals:  make(map[ast.Expr]int),
	}
}
