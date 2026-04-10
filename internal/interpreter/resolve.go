package interpreter

import (
	"loxlang/internal/ast"
	"loxlang/internal/tokenizer"
)

func (i *Interpreter) Resolve(expr ast.Expr, depth int) {
	i.locals[expr] = depth
}

func (i *Interpreter) lookUpVariable(name tokenizer.Token, expr ast.Expr) (any, bool) {
	distance, ok := i.locals[expr]
	if !ok {
		return i.globals.Get(name.Lexeme())
	}
	return i.env.GetAt(distance, name.Lexeme())
}
