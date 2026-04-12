package tokenizer

func (t *Tokenizer) Tokenize() []Token {
	for !t.isAtEnd() {
		t.start = t.current
		t.scanToken()
	}

	t.tokens = append(t.tokens, Token{
		tokenType: EOF,
		lexeme:    "",
		literal:   nil,
		line:      t.line,
	})

	return t.tokens
}
