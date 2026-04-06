package astprinter

import (
	"fmt"
	"loxlang/internal/ast"
	"strings"
)

type AstPrinter struct {
}

func New() *AstPrinter {
	return &AstPrinter{}
}

func (p *AstPrinter) Print(statements []ast.Stmt) string {
	var b strings.Builder

	for _, statement := range statements {
		b.WriteString(fmt.Sprintf("%v\n", statement.Accept(p)))
	}

	return b.String()
}
