package parser

import (
	"loxlang/internal/ast"
	"loxlang/internal/tokenizer"
	"loxlang/pkg/constants"
	"loxlang/pkg/loxerror"
)

type StmtParser interface {
	declaration() ast.Stmt
	funDecl() ast.Stmt
	function() ast.Stmt
	varDecl() ast.Stmt
	statement() ast.Stmt
	exprStatement() ast.Stmt
	forStatement() ast.Stmt
	ifStatement() ast.Stmt
	printStatement() ast.Stmt
	whileStatement() ast.Stmt
	block() []ast.Stmt
}

// declaration -> funDecl | varDecl | statement ;
func (p *Parser) declaration() ast.Stmt {
	if p.match(tokenizer.FUN) {
		return p.funDecl()
	}
	if p.match(tokenizer.VAR) {
		return p.varDecl()
	}
	return p.statement()
}

// funDecl -> "fun" function ;
func (p *Parser) funDecl() ast.Stmt {
	return p.function(constants.Function)
}

// function -> IDENTIFIER "(" parameters? ")" block ;
func (p *Parser) function(kind constants.CallableKind) ast.Stmt {
	name := p.consume(tokenizer.IDENTIFIER, loxerror.ErrExpectedIdenAfterFun)

	var params []tokenizer.Token

	p.consume(tokenizer.LEFT_PAREN, loxerror.ErrExpectedOpeningParenAfterCallableKindIden(kind))
	if !p.check(tokenizer.RIGHT_PAREN) {
		params = p.parameters()
	}
	p.consume(tokenizer.RIGHT_PAREN, loxerror.ErrExpectedClosingParenAfterParams)

	p.consume(tokenizer.LEFT_BRACE, loxerror.ErrExpectedOpeningBraceBeforeCallableKindBody(kind))

	body := p.block()

	return ast.NewFunction(*name, params, body)
}

// parameters -> IDENTIFIER ( "," IDENTIFIER )* ;
func (p *Parser) parameters() []tokenizer.Token {
	params := make([]tokenizer.Token, 0)
	params = append(params, *p.consume(tokenizer.IDENTIFIER, loxerror.ErrExpectParameterName))

	for p.match(tokenizer.COMMA) {
		if len(params) > constants.MAX_ARGUMENTS_COUNT {
			p.report(p.peek(), loxerror.ErrMaximumParamsCountExceeded)
		}
		params = append(params, *p.consume(tokenizer.IDENTIFIER, loxerror.ErrExpectParameterName))
	}
	return params
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

// statement -> exprStatement | forStatement | ifStatement | printStatement | whileStatement | block ;
func (p *Parser) statement() ast.Stmt {
	if p.match(tokenizer.FOR) {
		return p.forStatement()
	}
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

// "for" "(" ( varDecl | exprStmt | ";" ) expression? ";" expression? ")" statement ;
func (p *Parser) forStatement() ast.Stmt {
	p.consume(tokenizer.LEFT_PAREN, loxerror.ErrExpectedOpeningParenAfterFor)

	var initializer ast.Stmt

	if p.match(tokenizer.SEMICOLON) {
		initializer = nil
	} else if p.match(tokenizer.VAR) {
		initializer = p.varDecl()
	} else {
		initializer = p.exprStatement()
	}

	var condition ast.Expr

	if !p.check(tokenizer.SEMICOLON) {
		condition = p.expression()
	}

	p.consume(tokenizer.SEMICOLON, loxerror.ErrExpectedSemicolonAfterExpression)

	var modifier ast.Expr

	if !p.check(tokenizer.RIGHT_PAREN) {
		modifier = p.expression()
	}

	p.consume(tokenizer.RIGHT_PAREN, loxerror.ErrExpectedClosingParenAfterExpression)

	body := p.statement()

	return ast.NewFor(initializer, condition, modifier, body)
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
