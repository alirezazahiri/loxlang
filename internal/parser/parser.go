package parser

import (
	"loxlang/internal/tokenizer"
	"loxlang/pkg/loxerror"
)

type Parser struct {
	tokens  []tokenizer.Token
	current int
	line    int
	errors  []*loxerror.LoxError
}

func New(tokens []tokenizer.Token) *Parser {
	return &Parser{
		tokens: tokens,
		line:   1,
		errors: []*loxerror.LoxError{},
	}
}

func (t *Parser) Errors() []*loxerror.LoxError {
	return t.errors
}

func (t *Parser) HasErrors() bool {
	return len(t.errors) > 0
}
