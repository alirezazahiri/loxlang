package interpreter

import (
	"loxlang/internal/ast"
)

// Resolve panics if there is a resolution error
func (i *Interpreter) Interpret(statements []ast.Stmt) {
	for _, statement := range statements {
		i.execute(statement)
	}
}
