package lox

import (
	"loxlang/internal/ast"
	"loxlang/internal/parser"
	"loxlang/internal/tokenizer"
)

func (l *LoxLang) Parse(tokens []tokenizer.Token) []ast.Stmt {
	p := parser.New(tokens)
	result := p.Parse()

	l.hadError = p.HasErrors()
	l.errors = p.Errors()

	return result
}
