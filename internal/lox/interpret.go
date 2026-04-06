package lox

import (
	"loxlang/internal/ast"
	"loxlang/internal/interpreter"
)

func (l *LoxLang) Interpret(statements []ast.Stmt) {
	i := interpreter.New()
	err := i.Interpret(statements)

	l.hadError = err != nil
	l.errors = append(l.errors, err)
}
