package tokenizer

import (
	"loxlang/pkg/loxerror"
)

type Tokenizer struct {
	content string
	tokens  []Token
	start   int
	current int
	line    int
	errors  []*loxerror.LoxError
}

func New(content string) *Tokenizer {
	return &Tokenizer{
		content: content,
		tokens:  []Token{},
		start:   0,
		current: 0,
		line:    1,
		errors:  []*loxerror.LoxError{},
	}
}

func (t *Tokenizer) Errors() []*loxerror.LoxError {
	return t.errors
}

func (t *Tokenizer) HasErrors() bool {
	return len(t.errors) > 0
}
