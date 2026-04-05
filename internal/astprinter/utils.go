package astprinter

import (
	"loxlang/internal/ast"
	"strings"
)

func (p *AstPrinter) parenthesize(name string, exprs ...ast.Expr) string {
	var b strings.Builder
	b.WriteString("(")
	b.WriteString(name)

	for _, expr := range exprs {
		b.WriteString(" ")
		b.WriteString(expr.Accept(p).(string))
	}

	b.WriteString(")")
	return b.String()
}
