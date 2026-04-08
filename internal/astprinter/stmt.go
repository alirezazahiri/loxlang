package astprinter

import (
	"fmt"
	"loxlang/internal/ast"
	"strings"
)

func (p *AstPrinter) VisitExpressionStmt(stmt *ast.Expression) any {
	return p.parenthesize("stmt", stmt.Expression)
}
func (p *AstPrinter) VisitPrintStmt(stmt *ast.Print) any {
	return p.parenthesize("print", stmt.Expression)
}
func (p *AstPrinter) VisitVarStmt(stmt *ast.Var) any {
	return p.parenthesize(fmt.Sprintf("var::%s", stmt.Name.Lexeme()), stmt.Initializer)
}
func (p *AstPrinter) VisitBlockStmt(stmt *ast.Block) any {
	var b strings.Builder

	b.WriteString("(block\n")

	for _, s := range stmt.Statements {
		b.WriteString(" ")
		b.WriteString(fmt.Sprint(s.Accept(p)))
	}

	b.WriteString("\n)")
	return b.String()
}

func (p *AstPrinter) VisitIfStmt(stmt *ast.If) any {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("(if %v", stmt.Condition.Accept(p)))
	b.WriteString(fmt.Sprintf("(then %v)", stmt.ThenBranch.Accept(p)))
	b.WriteString(fmt.Sprintf("(else %v)", stmt.ElseBranch.Accept(p)))
	b.WriteString(")")
	return b.String()
}
