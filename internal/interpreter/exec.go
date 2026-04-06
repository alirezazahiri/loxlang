package interpreter

import (
	"loxlang/internal/ast"
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
