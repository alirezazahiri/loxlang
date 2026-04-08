package parser

import (
	"loxlang/internal/ast"
	"loxlang/internal/tokenizer"
	"loxlang/pkg/loxerror"
)

type StmtParser interface {
	declaration() ast.Stmt
	statement() ast.Stmt
	varDecl() ast.Stmt
	exprStatement() ast.Stmt
	printStatement() ast.Stmt
	block() ast.Stmt
}

// declaration -> varDecl | statement ;
func (p *Parser) declaration() ast.Stmt {
	if p.match(tokenizer.VAR) {
		return p.varDecl()
	}
	return p.statement()
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

// statement -> exprStatement | ifStatement | printStatement | whileStatement | block ;
func (p *Parser) statement() ast.Stmt {
	if p.match(tokenizer.WHILE) {
		return p.whileStatement()
	}
	if p.match(tokenizer.IF) {
		return p.ifStatement()
	}
	if p.match(tokenizer.PRINT) {
		return p.printStatement()
	}
	if p.match(tokenizer.LEFT_BRACE) {
		return ast.NewBlock(p.block())
	}
	return p.exprStatement()
}

// exprStatement -> expression ";" ;
func (p *Parser) exprStatement() ast.Stmt {
	expr := p.expression()
	p.consume(tokenizer.SEMICOLON, loxerror.ErrExpectedSemicolonAfterExpression)
	return ast.NewExpression(expr)
}

// "if" "(" expression ")" statement ( "else" statement )? ;
func (p *Parser) ifStatement() ast.Stmt {
	p.consume(tokenizer.LEFT_PAREN, loxerror.ErrExpectedOpeningParenAfterIf)
	expr := p.expression()
	p.consume(tokenizer.RIGHT_PAREN, loxerror.ErrExpectedClosingParenAfterExpression)
	thenBlock := p.statement()

	if p.match(tokenizer.ELSE) {
		return ast.NewIf(expr, thenBlock, p.statement())
	}

	return ast.NewIf(expr, thenBlock, nil)
}

// printStatement -> "print" expression ";" ;
func (p *Parser) printStatement() ast.Stmt {
	value := p.expression()
	p.consume(tokenizer.SEMICOLON, loxerror.ErrExpectedSemicolonAfterExpression)
	return ast.NewPrint(value)
}

// whileStatement -> "while" "(" expression ")" statement ;
func (p *Parser) whileStatement() ast.Stmt {
	p.consume(tokenizer.LEFT_PAREN, loxerror.ErrExpectedOpeningParenAfterWhile)
	cond := p.expression()
	p.consume(tokenizer.RIGHT_PAREN, loxerror.ErrExpectedClosingParenAfterExpression)
	stmt := p.statement()

	return ast.NewWhile(cond, stmt)
}

// block -> "{" declaration "}" ;
func (p *Parser) block() []ast.Stmt {
	statements := make([]ast.Stmt, 0)

	for !p.check(tokenizer.RIGHT_BRACE) && !p.isAtEnd() {
		statements = append(statements, p.declaration())
	}

	p.consume(tokenizer.RIGHT_BRACE, loxerror.ErrUnexpectedUnclosedBlock)

	return statements
}
