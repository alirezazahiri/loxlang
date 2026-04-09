package tokenizer

import "fmt"

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   any
	line      int
}

func NewToken(t TokenType, lexeme string, literal any, line int) Token {
	return Token{
		tokenType: t,
		lexeme:    lexeme,
		literal:   literal,
		line:      line,
	}
}

func (t *Token) Type() TokenType {
	return t.tokenType
}

func (t *Token) Lexeme() string {
	return t.lexeme
}

func (t *Token) Literal() any {
	return t.literal
}

func (t *Token) Line() int {
	return t.line
}

func (t *Token) ToString() string {
	return fmt.Sprintf("type=%d\tlexeme=%s\tliteral=%v", t.tokenType, t.lexeme, t.literal)
}
