package lox

import (
	"loxlang/internal/environment"
	"loxlang/pkg/loxerror"
)

type LoxLang struct {
	env      *environment.Environment
	hadError bool
	errors   []*loxerror.LoxError
}

func New() *LoxLang {
	return &LoxLang{
		env:      environment.New(nil),
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

	statements := l.Parse(tokens)
	if l.hadError {
		return l.errors
	}

	l.Interpret(statements)

	if l.hadError {
		return l.errors
	}

	return nil
}

func (l *LoxLang) HasErrors() bool {
	return l.hadError
}
