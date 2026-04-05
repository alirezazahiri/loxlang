package lox

import (
	"loxlang/internal/tokenizer"
)

func (l *LoxLang) Tokenize(src string) []tokenizer.Token {
	t := tokenizer.New(src)
	tokens := t.Tokenize()

	l.hadError = t.HasErrors()
	l.errors = t.Errors()

	return tokens
}
