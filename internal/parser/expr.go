package parser

import (
	"loxlang/internal/ast"
	"loxlang/internal/tokenizer"
	"loxlang/pkg/loxerror"
)

type ExprParser interface {
	expression() ast.Expr
	// binary
	assignment() ast.Expr
	equality() ast.Expr
	comparison() ast.Expr
	term() ast.Expr
	factor() ast.Expr
	// unary
	unary() ast.Expr
	// grouping, literal
	primary() ast.Expr
}

// expression -> assignment :
func (p *Parser) expression() ast.Expr {
	return p.assignment()
}

// assignment -> IDENTIFIER "=" assignment | logic_or
func (p *Parser) assignment() ast.Expr {
	expr := p.logicOr()
	if p.match(tokenizer.EQUAL) {
		equals := p.previous()
		value := p.assignment()

		if variable, ok := expr.(*ast.Variable); ok {
			name := variable.Name
			return ast.NewAssign(name, value)
		}

		p.report(equals, loxerror.ErrUnexpectedAssignmentExpression)
	}

	return expr
}

// logic_or -> logic_and ( "or" logic_and )*
func (p *Parser) logicOr() ast.Expr {
	expr := p.logicAnd()

	for p.match(tokenizer.OR) {
		operator := *p.previous()
		right := p.logicAnd()
		expr = ast.NewLogical(expr, operator, right)
	}

	return expr
}

// logic_and -> equality ( "and" equality )*
func (p *Parser) logicAnd() ast.Expr {
	expr := p.equality()

	for p.match(tokenizer.AND) {
		operator := *p.previous()
		right := p.equality()
		expr = ast.NewLogical(expr, operator, right)
	}

	return expr
}

// equality -> comparison ( ( "!=" | "==" ) comparison )* ;
func (p *Parser) equality() ast.Expr {
	expr := p.comparison()
	for p.match(tokenizer.BANG_EQUAL, tokenizer.EQUAL_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expr = ast.NewBinary(expr, *operator, right)
	}
	return expr
}

// comparison -> term ( ( ">" | ">=" | "<" | "<=" ) term )* ;
func (p *Parser) comparison() ast.Expr {
	expr := p.term()
	for p.match(tokenizer.GREATER, tokenizer.GREATER_EQUAL, tokenizer.LESS, tokenizer.LESS_EQUAL) {
		operator := p.previous()
		right := p.term()
		expr = ast.NewBinary(expr, *operator, right)
	}
	return expr
}

// term -> factor ( ( "-" | "+" ) factor )* ;
func (p *Parser) term() ast.Expr {
	expr := p.factor()
	for p.match(tokenizer.MINUS, tokenizer.PLUS) {
		operator := p.previous()
		right := p.factor()
		expr = ast.NewBinary(expr, *operator, right)
	}
	return expr

}

// factor -> unary ( ( "/" | "*" ) unary )* ;
func (p *Parser) factor() ast.Expr {
	expr := p.unary()
	for p.match(tokenizer.SLASH, tokenizer.STAR) {
		operator := p.previous()
		right := p.unary()
		expr = ast.NewBinary(expr, *operator, right)
	}
	return expr
}

// unary -> ( "!" | "-" ) unary | primary ;
func (p *Parser) unary() ast.Expr {
	if p.match(tokenizer.BANG, tokenizer.MINUS) {
		operator := p.previous()
		right := p.unary()
		return ast.NewUnary(*operator, right)
	}
	expr := p.primary()
	return expr
}

// primary -> NUMBER | STRING | "true" | "false" | "nil" | "(" expression ")" ;
func (p *Parser) primary() ast.Expr {
	if p.match(tokenizer.NUMBER, tokenizer.STRING, tokenizer.TRUE, tokenizer.FALSE, tokenizer.NIL) {
		return ast.NewLiteral(p.previous().Literal())
	}

	if p.match(tokenizer.IDENTIFIER) {
		return ast.NewVariable(*p.previous())
	}

	if p.match(tokenizer.LEFT_PAREN) {
		expr := p.expression()

		token := p.consume(tokenizer.RIGHT_PAREN, loxerror.ErrUnexpectedToken(p.peek().Lexeme(), ")"))

		if token != nil {
			return ast.NewGrouping(expr)
		}
	}

	p.report(p.peek(), loxerror.ErrUnexpectedTokenAfterExpression)

	return nil
}
