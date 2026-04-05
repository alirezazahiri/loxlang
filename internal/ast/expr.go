// THIS FILE IS AUTO GENERATED AND IS READ-ONLY (DO NOT MODIFY IT)

package ast

import "loxlang/internal/tokenizer"

type Expr interface {
	Accept(visitor ExprVisitor) any
}

type Binary struct {
	Left     Expr
	Operator tokenizer.Token
	Right    Expr
}

func NewBinary(left Expr, operator tokenizer.Token, right Expr) *Binary {
	return &Binary{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}

func (e *Binary) Accept(visitor ExprVisitor) any {
	return visitor.VisitBinaryExpr(e)
}

type Grouping struct {
	Expression Expr
}

func NewGrouping(expression Expr) *Grouping {
	return &Grouping{
		Expression: expression,
	}
}

func (e *Grouping) Accept(visitor ExprVisitor) any {
	return visitor.VisitGroupingExpr(e)
}

type Literal struct {
	Value any
}

func NewLiteral(value any) *Literal {
	return &Literal{
		Value: value,
	}
}

func (e *Literal) Accept(visitor ExprVisitor) any {
	return visitor.VisitLiteralExpr(e)
}

type Unary struct {
	Operator tokenizer.Token
	Right    Expr
}

func NewUnary(operator tokenizer.Token, right Expr) *Unary {
	return &Unary{
		Operator: operator,
		Right:    right,
	}
}

func (e *Unary) Accept(visitor ExprVisitor) any {
	return visitor.VisitUnaryExpr(e)
}

type ExprVisitor interface {
	VisitBinaryExpr(expr *Binary) any
	VisitGroupingExpr(expr *Grouping) any
	VisitLiteralExpr(expr *Literal) any
	VisitUnaryExpr(expr *Unary) any
}
