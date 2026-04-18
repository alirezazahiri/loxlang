package lox

import (
	"loxlang/internal/builtin"
	builtinclasses "loxlang/internal/builtin/classes"
	"loxlang/internal/environment"
	"loxlang/pkg/constants"
	"loxlang/pkg/loxerror"
)

type LoxLang struct {
	globals  *environment.Environment
	hadError bool
	errors   []*loxerror.LoxError
}

func New() *LoxLang {
	return &LoxLang{
		hadError: false,
		errors:   nil,
		globals:  environment.New(nil),
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

	l.registerBuiltins()
	l.Interpret(statements)

	if l.hadError {
		return l.errors
	}

	return nil
}

func (l *LoxLang) HasErrors() bool {
	return l.hadError
}

func (l *LoxLang) Errors() []*loxerror.LoxError {
	return l.errors
}

func (l *LoxLang) registerBuiltins() {
	l.globals.Define(constants.CLOCK_FN, builtin.NewClockFn())
	l.globals.Define(constants.INPUT_FN, builtin.NewInputFn())
	l.registerBuiltinClasses()
}

func (l *LoxLang) registerBuiltinClasses() {
	l.globals.Define(constants.CONSOLE_CLASS_INSTANCE, builtinclasses.Register(builtinclasses.NewConsoleClass))
}
