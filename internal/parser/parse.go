package parser

import (
	"loxlang/internal/ast"
)

func (p *Parser) Parse() []ast.Stmt {
	statements := make([]ast.Stmt, 0)

	for !p.isAtEnd() {
		defer func() {
			if r := recover(); r != nil {
				p.synchronize()
			}
		}()

		statements = append(statements, p.declaration())
	}

	return statements
}
