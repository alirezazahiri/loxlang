package interpreter

import (
	"loxlang/internal/ast"
)

func (i *Interpreter) Interpret(statements []ast.Stmt) {
	for _, statement := range statements {
		i.execute(statement)
	}
}
