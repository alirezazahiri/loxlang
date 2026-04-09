package interpreter

import "loxlang/internal/ast"

func (i *Interpreter) evaluate(expr ast.Expr) any {
	return expr.Accept(i)
}
