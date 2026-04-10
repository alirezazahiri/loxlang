package lox

import (
	"loxlang/internal/ast"
	"loxlang/internal/interpreter"
	"loxlang/pkg/loxerror"
)

func (l *LoxLang) Interpret(statements []ast.Stmt) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(*loxerror.LoxError); ok {
				l.hadError = e != nil
				l.errors = append(l.errors, e)
			} else {
				panic(r)
			}
		}
	}()

	i := interpreter.New(l.globals)

	l.Resolve(i, statements)

	i.Interpret(statements)
}
