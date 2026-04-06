package interpreter

import (
	"loxlang/internal/ast"
	"loxlang/pkg/loxerror"
)

func (i *Interpreter) Interpret(statements []ast.Stmt) *loxerror.LoxError {
	for _, statement := range statements {
		if err := i.execute(statement); err != nil {
			return err
		}
	}

	return nil
}
