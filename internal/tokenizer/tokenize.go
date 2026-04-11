package tokenizer

import "fmt"

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

	for _, tok := range t.tokens {
		fmt.Println(tok)
	}

	return t.tokens
}
