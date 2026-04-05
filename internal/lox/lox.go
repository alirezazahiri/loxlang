package lox

import (
	"loxlang/pkg/loxerror"
)

type LoxLang struct {
	hadError bool
	hadRuntimeError bool
	errors   []*loxerror.LoxError
}

func New() *LoxLang {
	return &LoxLang{
		hadError: false,
		errors:   nil,
	}
}

func (l *LoxLang) Run(code string) []*loxerror.LoxError {
	l.errors = nil
	l.hadError = false

	tokens := l.Tokenize(code)
	if l.hadError {
		return l.errors
	}

	expr := l.Parse(tokens)
	if l.hadError {
		return l.errors
	}

	l.Interpret(expr)

	if l.hadError {
		return l.errors
	}

	return nil
}


func (l *LoxLang) HasErrors() bool {
	return l.hadError
}
