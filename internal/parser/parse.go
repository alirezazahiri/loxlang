package parser

import (
	"loxlang/internal/ast"
)

func (p *Parser) Parse() ast.Expr {
	return p.expression()
}
