package lox

import (
	"loxlang/internal/ast"
	"loxlang/internal/interpreter"
	"loxlang/pkg/loxerror"
)

func (l *LoxLang) Interpret(st ast.Expr) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(*loxerror.LoxError); ok {
				l.hadError = true
				l.errors = append(l.errors, err)
			} else {
				panic(r)
			}
		}
	}()

	i := interpreter.New()
	i.Interpret(st)
}
