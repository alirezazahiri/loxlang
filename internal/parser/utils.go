package parser

import (
	"loxlang/internal/tokenizer"
	"loxlang/pkg/loxerror"
)

func (p *Parser) match(types ...tokenizer.TokenType) bool {
	for _, tokenType := range types {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(t tokenizer.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type() == t
}

func (p *Parser) peek() *tokenizer.Token {
	return &p.tokens[p.current]
}

func (p *Parser) previous() *tokenizer.Token {
	return &p.tokens[p.current-1]
}

func (p *Parser) advance() *tokenizer.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type() == tokenizer.EOF
}

func (p *Parser) consume(tokenType tokenizer.TokenType, message loxerror.LoxErrorMessage) *tokenizer.Token {
	if p.check(tokenType) {
		return p.advance()
	}

	p.report(p.peek(), message)

	return nil
}

func (p *Parser) report(token *tokenizer.Token, message loxerror.LoxErrorMessage) {
	var err *loxerror.LoxError = nil

	line := token.Line()
	where := token.Lexeme()

	if token.Type() == tokenizer.EOF {
		err = loxerror.NewParserError(line, "end", message)
	} else {
		err = loxerror.NewParserError(line, where, message)
	}

	p.errors = append(p.errors, err)
}

func (p *Parser) synchronize() {
	p.advance()

	for !p.isAtEnd() {
		if p.previous().Type() == tokenizer.SEMICOLON {
			return
		}

		switch p.peek().Type() {
		case tokenizer.CLASS:
		case tokenizer.FUN:
		case tokenizer.VAR:
		case tokenizer.FOR:
		case tokenizer.IF:
		case tokenizer.WHILE:
		case tokenizer.PRINT:
		case tokenizer.RETURN:
			return
		}
	}

	p.advance()
}
