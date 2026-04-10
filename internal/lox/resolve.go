package lox

import (
	"loxlang/internal/ast"
	"loxlang/internal/interpreter"
	"loxlang/internal/resolver"
)

func (l *LoxLang) Resolve(interpreter *interpreter.Interpreter, statements []ast.Stmt) {
	r := resolver.New(interpreter)
	r.Resolve(statements)
}