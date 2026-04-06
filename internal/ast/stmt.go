// THIS FILE IS AUTO GENERATED AND IS READ-ONLY (DO NOT MODIFY IT)

package ast

import (
	"loxlang/internal/tokenizer"
)

type Stmt interface {
	Accept(visitor StmtVisitor) any
}

type Expression struct {
	Expression Expr
}

func NewExpression(expression Expr) *Expression {
	return &Expression{
		Expression: expression,
	}
}

func (e *Expression) Accept(visitor StmtVisitor) any {
	return visitor.VisitExpressionStmt(e)
}

type Print struct {
	Expression Expr
}

func NewPrint(expression Expr) *Print {
	return &Print{
		Expression: expression,
	}
}

func (e *Print) Accept(visitor StmtVisitor) any {
	return visitor.VisitPrintStmt(e)
}

type Var struct {
	Name        tokenizer.Token
	Initializer Expr
}

func NewVar(name tokenizer.Token, initializer Expr) *Var {
	return &Var{
		Name:        name,
		Initializer: initializer,
	}
}

func (e *Var) Accept(visitor StmtVisitor) any {
	return visitor.VisitVarStmt(e)
}

type StmtVisitor interface {
	VisitExpressionStmt(stmt *Expression) any
	VisitPrintStmt(stmt *Print) any
	VisitVarStmt(stmt *Var) any
}
