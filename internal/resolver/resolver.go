package resolver

import (
	"loxlang/internal/interpreter"
	"loxlang/internal/scope"
	"loxlang/internal/tokenizer"
	"loxlang/pkg/constants"
)

type Resolver struct {
	interpreter     *interpreter.Interpreter
	scopes          scope.Scopes
	currentFunction constants.CallableKind
	currentClass    constants.ClassKind
	initializing    *tokenizer.Token
}

func New(interpreter *interpreter.Interpreter) *Resolver {
	resolver := &Resolver{
		interpreter:     interpreter,
		scopes:          make(scope.Scopes, 0),
		currentFunction: constants.CallableKindNone,
		currentClass:    constants.ClassKindNone,
	}

	// start global scope
	resolver.beginScope()

	return resolver
}
