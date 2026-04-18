package parser

import (
	"loxlang/internal/ast"
	"loxlang/internal/tokenizer"
	"loxlang/pkg/constants"
	"loxlang/pkg/loxerror"
)

type StmtParser interface {
	declaration() ast.Stmt
	classDecl() ast.Stmt
	funDecl() ast.Stmt
	function() ast.Stmt
	setter() ast.Stmt
	parameters() ast.Stmt
	varDecl() ast.Stmt
	statement() ast.Stmt
	exprStatement() ast.Stmt
	forStatement() ast.Stmt
	ifStatement() ast.Stmt
	returnStatement() ast.Stmt
	printStatement() ast.Stmt
	whileStatement() ast.Stmt
	block() []ast.Stmt
}

// declaration -> classDel | funDecl | varDecl | statement ;
func (p *Parser) declaration() ast.Stmt {
	if p.match(tokenizer.CLASS) {
		return p.classDecl()
	}
	if p.match(tokenizer.FUN) {
		return p.funDecl()
	}
	if p.match(tokenizer.VAR) {
		return p.varDecl()
	}
	return p.statement()
}

// classDecl -> "class" IDENTIFIER ( "<" IDENTIFIER )? "{" ( function | setter )* "}" ;
func (p *Parser) classDecl() ast.Stmt {
	className := p.consume(tokenizer.IDENTIFIER, loxerror.ErrExpectedIdenForClass)

	var superClass *ast.Variable

	if p.match(tokenizer.LESS) {
		superClass = ast.NewVariable(*p.consume(tokenizer.IDENTIFIER, loxerror.ErrExpectedIdenForSuperClass))
	}

	p.consume(tokenizer.LEFT_BRACE, loxerror.ErrExpectedOpeningBraceBeforeClassBody)

	methods := make([]ast.Function, 0)
	for p.check(tokenizer.IDENTIFIER) || p.check(tokenizer.SET) {
		switch p.peek().Type() {
		case tokenizer.SET:
			methods = append(methods, *p.setter())
		case tokenizer.IDENTIFIER:
			methods = append(methods, *p.function(constants.CallableKindSetter))
		}
	}
	p.consume(tokenizer.RIGHT_BRACE, loxerror.ErrExpectedClosingBraceAfterClassBody)

	return ast.NewClass(*className, superClass, methods)
}

// funDecl -> "fun" function ;
func (p *Parser) funDecl() ast.Stmt {
	return p.function(constants.CallableKindFunction)
}

// function -> "set"? IDENTIFIER ( "(" parameters? ")" )? block ;
func (p *Parser) function(kind constants.CallableKind) *ast.Function {
	name := p.consume(tokenizer.IDENTIFIER, loxerror.ErrExpectedIdenAfterFun)
	if p.match(tokenizer.LEFT_PAREN) {
		params, body := p.extractParamsAndBody(kind)
		return ast.NewFunction(*name, params, body, false, false)
	} else if p.match(tokenizer.LEFT_BRACE) {
		name.SetLexeme(constants.GETTER_PREFIX + name.Lexeme())
		return ast.NewFunction(*name, nil, p.block(), true, false)
	}
	if !p.match(tokenizer.LEFT_PAREN, tokenizer.LEFT_BRACE) {
		p.report(p.peek(), loxerror.ErrExpectedOpeningParenAfterCallableKindIden(kind))
	}
	return nil
}

// setter -> "set" IDENTIFIER "(" param ")" block ;
func (p *Parser) setter() *ast.Function {
	p.consume(tokenizer.SET, loxerror.ErrExpectedSetKeywordBeforeSet)
	name := p.consume(tokenizer.IDENTIFIER, loxerror.ErrExpectedIdenAfterSet)
	p.consume(tokenizer.LEFT_PAREN, loxerror.ErrExpectedOpeningParenAfterCallableKindIden(constants.CallableKindSetter))
	params, body := p.extractParamsAndBody(constants.CallableKindSetter)
	if len(params) != 1 {
		p.report(name, loxerror.ErrSettersExpectExactlyOneParam)
	}
	name.SetLexeme(constants.SETTER_PREFIX + name.Lexeme())
	return ast.NewFunction(*name, params, body, false, true)
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

// statement -> exprStatement | forStatement | ifStatement | returnStatement | printStatement | whileStatement | block ;
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
	if p.match(tokenizer.RETURN) {
		return p.returnStatement()
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

// forStmt -> "for" "(" ( varDecl | exprStmt | ";" ) expression? ";" expression? ")" statement ;
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

// ifStmt -> "if" "(" expression ")" statement ( "else" statement )? ;
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

// return -> "return" expression? ";" ;
func (p *Parser) returnStatement() ast.Stmt {
	keyword := *p.previous()
	var value ast.Expr = nil
	if !p.check(tokenizer.SEMICOLON) {
		value = p.expression()
	}
	p.consume(tokenizer.SEMICOLON, loxerror.ErrExpectedSemicolonAfterExpression)
	return ast.NewReturn(keyword, value)
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
