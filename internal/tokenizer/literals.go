package tokenizer

import (
	"loxlang/pkg/loxerror"
	"loxlang/pkg/utils"
	"strconv"
	"strings"
)

func (t *Tokenizer) string() {
	for t.peek() != '"' && !t.isAtEnd() {
		if t.peek() == '\n' {
			t.line++
		}
		t.advance()
	}

	if t.isAtEnd() {
		raw := t.content[t.start:t.current] // skip opening "
		before, _, found := strings.Cut(raw, "\n")
		s := ""
		if found {
			s = before
		} else {
			s = raw
		}
		s = strings.TrimSpace(s)
		err := loxerror.NewTokenizerError(t.line, s, loxerror.ErrUnterminatedString)
		t.errors = append(t.errors, err)
		return
	}

	t.advance()

	value := t.content[t.start+1 : t.current-1]
	t.addLiteral(STRING, value)
}

func (t *Tokenizer) number() {
	for utils.IsDigit(t.peek()) {
		t.advance()
	}

	// fractional part
	if t.peek() == '.' && utils.IsDigit(t.peekNext()) {
		t.advance() // '.'
		for utils.IsDigit(t.peek()) {
			t.advance()
		}
	}

	val, _ := strconv.ParseFloat(t.content[t.start:t.current], 64)
	t.addLiteral(NUMBER, val)
}
