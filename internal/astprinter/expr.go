package astprinter

import (
	"fmt"
	"loxlang/internal/ast"
	"strings"
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

func (p *AstPrinter) VisitAssignExpr(expr *ast.Assign) any {
	return p.parenthesize(fmt.Sprintf("assign::%v", expr.Name.Lexeme()), expr.Value)
}

func (p *AstPrinter) VisitLogicalExpr(expr *ast.Logical) any {
	return p.parenthesize(expr.Operator.Lexeme(), expr.Left, expr.Right)
}

func (p *AstPrinter) VisitCallExpr(expr *ast.Call) any {
	var b strings.Builder

	b.WriteString(p.parenthesize("call", expr.Callee))
	if len(expr.Arguments) > 0 {
		b.WriteString(p.parenthesize("arg", expr.Arguments...))
	} else {
		b.WriteString("(no-args)")
	}

	return b.String()
}


func (p *AstPrinter) VisitLambdaExpr(expr *ast.Lambda) any {
		var b strings.Builder
	
		b.WriteString("(anonymous-fun\n")
		params := expr.Params
		body := expr.Body
	
		paramsCount := 0
		if params != nil {
			paramsCount = len(params)
		}
		
		b.WriteString(fmt.Sprintf("(params::%d ", paramsCount))
		b.WriteString("anonymous-fun")
	
		for _, s := range body {
			b.WriteString(" ")
			b.WriteString(fmt.Sprint(s.Accept(p)))
		}
	
		b.WriteString("\n)")
		return b.String()
	
}