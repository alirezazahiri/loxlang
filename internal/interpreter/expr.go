package interpreter

import (
	"loxlang/internal/ast"
	"loxlang/internal/tokenizer"
	"loxlang/pkg/loxerror"
)

func (i *Interpreter) VisitLiteralExpr(expr *ast.Literal) any {
	return expr.Value
}

func (i *Interpreter) VisitBinaryExpr(expr *ast.Binary) any {
	left := i.evaluate(expr.Left)
	operator := expr.Operator
	right := i.evaluate(expr.Right)

	switch operator.Type() {
	case tokenizer.PLUS:
		i.checkMatchingOperandTypes(operator, left, right)

		ln, leftIsNumber := left.(float64)
		rn, rightIsNumber := right.(float64)

		if leftIsNumber && rightIsNumber {
			return ln + rn
		}

		ls, leftIsString := left.(string)
		rs, rightIsString := right.(string)

		if leftIsString && rightIsString {
			return ls + rs
		}

		break
	case tokenizer.MINUS:
		i.checkNumberOperands(operator, left, right)
		return left.(float64) - right.(float64)
	case tokenizer.STAR:
		i.checkNumberOperands(operator, left, right)
		return left.(float64) * right.(float64)
	case tokenizer.SLASH:
		i.checkNumberOperands(operator, left, right)
		return left.(float64) / right.(float64)
	case tokenizer.GREATER:
		i.checkNumberOperands(operator, left, right)
		return left.(float64) > right.(float64)
	case tokenizer.GREATER_EQUAL:
		i.checkNumberOperands(operator, left, right)
		return left.(float64) >= right.(float64)
	case tokenizer.EQUAL_EQUAL:
		return i.isEqual(left, right)
	case tokenizer.BANG_EQUAL:
		return !i.isEqual(left, right)
	case tokenizer.LESS:
		i.checkNumberOperands(operator, left, right)
		return left.(float64) < right.(float64)
	case tokenizer.LESS_EQUAL:
		i.checkNumberOperands(operator, left, right)
		return left.(float64) <= right.(float64)
	}

	return nil
}

func (i *Interpreter) VisitUnaryExpr(expr *ast.Unary) any {
	operator := expr.Operator
	right := i.evaluate(expr.Right)

	switch operator.Type() {
	case tokenizer.BANG:
		return !i.isTruthy(right)
	case tokenizer.MINUS:
		i.checkNumberOperand(operator, right)
		return -right.(float64)
	}

	return nil
}

func (i *Interpreter) VisitGroupingExpr(expr *ast.Grouping) any {
	return i.evaluate(expr.Expression)
}

func (i *Interpreter) VisitVariableExpr(expr *ast.Variable) any {
	value, ok := i.env.Get(expr.Name.Lexeme())
	if !ok {
		i.report(loxerror.NewRuntimeError(expr.Name.Line(), expr.Name.Lexeme(), loxerror.ErrUndefinedVariable))
	}
	return value
}

func (i *Interpreter) VisitAssignExpr(expr *ast.Assign) any {
	value := i.evaluate(expr.Value)

	ok := i.env.Assign(expr.Name.Lexeme(), value)

	if !ok {
		i.report(loxerror.NewRuntimeError(expr.Name.Line(), expr.Name.Lexeme(), loxerror.ErrCannotAssignToNonDeclaredVariables))
	}

	return nil
}

func (i *Interpreter) VisitLogicalExpr(expr *ast.Logical) any {
	left := i.evaluate(expr.Left)

	switch expr.Operator.Type() {
	case tokenizer.OR:
		if i.isTruthy(left) {
			return left
		}
	case tokenizer.AND:
		if !i.isTruthy(left) {
			return left
		}
	}

	return i.evaluate(expr.Right)
}
