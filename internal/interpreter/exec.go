package interpreter

import (
	"loxlang/internal/ast"
	"loxlang/internal/environment"
	"loxlang/pkg/loxerror"
)

func (i *Interpreter) execute(statement ast.Stmt) (err *loxerror.LoxError) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(*loxerror.LoxError); ok {
				err = e
			} else {
				panic(r)
			}
		}
	}()
	statement.Accept(i)
	return nil
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
