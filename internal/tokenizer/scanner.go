package tokenizer

import (
	"loxlang/pkg/loxerror"
)

func (t *Tokenizer) scanToken() {
	c := t.advance()

	switch c {
	case '(':
		t.add(LEFT_PAREN)
	case ')':
		t.add(RIGHT_PAREN)
	case '{':
		t.add(LEFT_BRACE)
	case '}':
		t.add(RIGHT_BRACE)
	case ',':
		t.add(COMMA)
	case '.':
		t.add(DOT)
	case '-':
		t.add(MINUS)
	case '+':
		t.add(PLUS)
	case ';':
		t.add(SEMICOLON)
	case '*':
		t.add(STAR)

	case '!':
		if t.match('=') {
			t.add(BANG_EQUAL)
		} else {
			t.add(BANG)
		}

	case '=':
		if t.match('=') {
			t.add(EQUAL_EQUAL)
		} else {
			t.add(EQUAL)
		}

	case '<':
		if t.match('=') {
			t.add(LESS_EQUAL)
		} else {
			t.add(LESS)
		}

	case '>':
		if t.match('=') {
			t.add(GREATER_EQUAL)
		} else {
			t.add(GREATER)
		}

	case '/':
		// comments
		if t.match('/') {
			for t.peek() != '\n' && !t.isAtEnd() {
				t.advance()
			}
		} else {
			t.add(SLASH)
		}

	case ' ':
	case '\r':
	case '\t':
		// ignore whitespace

	case '\n':
		t.line++

	case '"':
		t.string()

	default:
		if isDigit(c) {
			t.number()
		} else if isAlpha(c) {
			t.identifier()
		} else {
			t.errors = append(t.errors, loxerror.NewTokenizerError(t.line, string(c), loxerror.ErrUnexpectedChar))
		}
	}
}

func (t *Tokenizer) add(tokenType TokenType) {
	text := t.content[t.start:t.current]
	t.tokens = append(t.tokens, NewToken(
		tokenType,
		text,
		nil,
		t.line,
	))
}

func (t *Tokenizer) addLiteral(tokenType TokenType, literal any) {
	text := t.content[t.start:t.current]
	t.tokens = append(t.tokens, NewToken(
		tokenType,
		text,
		literal,
		t.line,
	))
}

func (t *Tokenizer) identifier() {
	for isAlphaNumeric(t.peek()) {
		t.advance()
	}

	text := t.content[t.start:t.current]

	if tok, ok := keywords[text]; ok {
		switch tok {
		case TRUE:
			t.addLiteral(tok, true)
		case FALSE:
			t.addLiteral(tok, false)
		default:
			t.add(tok)
		}
	} else {
		t.add(IDENTIFIER)
	}
}
