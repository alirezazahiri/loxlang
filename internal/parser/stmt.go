package parser

import (
	"loxlang/internal/ast"
	"loxlang/internal/tokenizer"
	"loxlang/pkg/loxerror"
)

type StmtParser interface {
	program() ast.Stmt
	statement() ast.Stmt
	exprStatement() ast.Stmt
	printStatement() ast.Stmt
}

// declaration -> varDecl | statement ;
func (p *Parser) declaration() ast.Stmt {
	if p.match(tokenizer.VAR) {
		return p.varDecl()
	}
	return p.statement()
}

// statement -> exprStatement | printStatement ;
func (p *Parser) statement() ast.Stmt {
	if p.match(tokenizer.PRINT) {
		return p.printStatement()
	}
	return p.exprStatement()
}

// exprStatement -> expression ";" ;
func (p *Parser) exprStatement() ast.Stmt {
	expr := p.expression()
	p.consume(tokenizer.SEMICOLON, loxerror.ErrExpectedSemicolonAfterExpression)
	return ast.NewExpression(expr)
}

// varDecl -> "var" IDENTIFIER ( "=" expression )? ";"
func (p *Parser) varDecl() ast.Stmt {
	iden := p.consume(tokenizer.IDENTIFIER, loxerror.ErrExpectedIdenAfterVar)
	
	var initializer ast.Expr
	if p.match(tokenizer.EQUAL) {
		initializer = p.expression()
	}
	
	p.consume(tokenizer.SEMICOLON, loxerror.ErrExpectedInitializationOrSemicolon)
	return ast.NewVar(*iden, initializer)
}

// printStatement -> "print" expression ";" ;
func (p *Parser) printStatement() ast.Stmt {
	value := p.expression()
	p.consume(tokenizer.SEMICOLON, loxerror.ErrExpectedSemicolonAfterExpression)
	return ast.NewPrint(value)
}
