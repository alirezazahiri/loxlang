package astprinter

import (
	"fmt"
	"loxlang/internal/ast"
)

func (p *AstPrinter) VisitBinaryExpr(expr *ast.Binary) any {
	return p.parenthesize(expr.Operator.Lexeme(), expr.Left, expr.Right)
}

func (p *AstPrinter) VisitGroupingExpr(expr *ast.Grouping) any {
	return p.parenthesize("group", expr.Expression)
}

func (p *AstPrinter) VisitLiteralExpr(expr *ast.Literal) any {
	if expr.Value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", expr.Value)
}

func (p *AstPrinter) VisitUnaryExpr(expr *ast.Unary) any {
	return p.parenthesize(expr.Operator.Lexeme(), expr.Right)
}

func (p *AstPrinter) VisitVariableExpr(expr *ast.Variable) any {
	return p.parenthesize(expr.Name.Lexeme())
}
