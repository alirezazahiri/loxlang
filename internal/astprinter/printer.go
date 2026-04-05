package astprinter

import (
	"fmt"
	"loxlang/internal/ast"
)

type AstPrinter struct {
}

func New() *AstPrinter {
	return &AstPrinter{}
}

func (p *AstPrinter) Print(expr ast.Expr) string {
	return fmt.Sprintf("%v", expr.Accept(p))
}
