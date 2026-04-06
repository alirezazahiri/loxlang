package astprinter

import (
	"fmt"
	"loxlang/internal/ast"
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