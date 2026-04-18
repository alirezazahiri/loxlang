package interpreter

import (
	"loxlang/internal/ast"
	"loxlang/internal/environment"
)

func (i *Interpreter) execute(statement ast.Stmt) {
	statement.Accept(i)
}

func (i *Interpreter) executeBlock(statements []ast.Stmt, env *environment.Environment) {
	previous := i.env
	i.env = env

	defer func() {
		i.env = previous
	}()

	for _, statement := range statements {
		i.execute(statement)
	}
}
