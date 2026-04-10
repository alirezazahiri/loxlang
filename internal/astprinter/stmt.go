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
	if stmt.Initializer != nil {
		return p.parenthesize(fmt.Sprintf("var::%s", stmt.Name.Lexeme()), stmt.Initializer)
	}
	return fmt.Sprintf("var::%s", stmt.Name.Lexeme())
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
	if stmt.ElseBranch != nil {
		b.WriteString(fmt.Sprintf("(else %v)", stmt.ElseBranch.Accept(p)))
	}
	b.WriteString(")")
	return b.String()
}

func (p *AstPrinter) VisitWhileStmt(stmt *ast.While) any {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("(while %v", stmt.Condition.Accept(p)))
	b.WriteString(fmt.Sprintf("(then %v)", stmt.Then.Accept(p)))
	b.WriteString(")")
	return b.String()
}

func (p *AstPrinter) VisitForStmt(stmt *ast.For) any {
	var b strings.Builder
	init := stmt.Initializer
	cond := stmt.Condition
	mod := stmt.Modifier
	b.WriteString("(for ")
	if init != nil {
		b.WriteString(fmt.Sprintf("init::%v; ", init.Accept(p)))
	}
	if cond != nil {
		b.WriteString(p.parenthesize("cond", cond))
	}
	b.WriteString("; ")
	if mod != nil {
		b.WriteString(p.parenthesize("mod", mod))
	}
	b.WriteString(fmt.Sprintf(" (body %v)", stmt.Body.Accept(p)))
	b.WriteString(")")
	return b.String()
}

func (p *AstPrinter) VisitFunctionStmt(stmt *ast.Function) any {
	var b strings.Builder

	name := stmt.Name
	b.WriteString(fmt.Sprintf("(fun::%s\n", name.Lexeme()))
	params := stmt.Params
	body := stmt.Body

	paramsCount := 0
	if params != nil {
		paramsCount = len(params)
	}
	
	b.WriteString(fmt.Sprintf("(params::%d ", paramsCount))
	b.WriteString(fmt.Sprintf("%s::block", name.Lexeme()))

	for _, s := range body {
		b.WriteString(" ")
		b.WriteString(fmt.Sprint(s.Accept(p)))
	}

	b.WriteString("\n)")
	return b.String()
}

func (p *AstPrinter) VisitReturnStmt(stmt *ast.Return) any {
	return p.parenthesize("return", stmt.Value)
}

func (p *AstPrinter) VisitClassStmt(stmt *ast.Class) any {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("(class::%s ", stmt.Name.Lexeme()))
	for _, method := range stmt.Methods {
		b.WriteString(method.Accept(p).(string))
	}
	b.WriteString(" )")
	return b.String()
}
