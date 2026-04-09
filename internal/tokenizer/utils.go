package tokenizer

func (t *Tokenizer) isAtEnd() bool {
	return t.current >= len(t.content)
}

func (t *Tokenizer) advance() byte {
	ch := t.content[t.current]
	t.current++
	return ch
}

func (t *Tokenizer) match(expected byte) bool {
	if t.isAtEnd() {
		return false
	}
	if t.content[t.current] != expected {
		return false
	}
	t.current++
	return true
}

func (t *Tokenizer) peek() byte {
	if t.isAtEnd() {
		return 0
	}
	return t.content[t.current]
}

func (t *Tokenizer) peekNext() byte {
	if t.current+1 >= len(t.content) {
		return 0
	}
	return t.content[t.current+1]
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		c == '_'
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}
