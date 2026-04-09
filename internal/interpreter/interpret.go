package interpreter

import (
	"loxlang/internal/ast"
	"loxlang/pkg/loxerror"
)

func (i *Interpreter) Interpret(statements []ast.Stmt) (err *loxerror.LoxError) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(*loxerror.LoxError); ok {
				err = e
			} else {
				panic(r)
			}
		}
	}()

	for _, statement := range statements {
		i.execute(statement)
	}

	return nil
}
