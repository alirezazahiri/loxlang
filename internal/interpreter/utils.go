package interpreter

import (
	"fmt"
	"loxlang/internal/ast"
	"loxlang/internal/tokenizer"
	"loxlang/pkg/loxerror"
	"strings"
)

func (i *Interpreter) isTruthy(value any) bool {
	if value == nil {
		return false
	}

	if b, ok := value.(bool); ok {
		return b
	}

	switch n := value.(type) {
	case float64:
		return n != 0
	}

	if s, ok := value.(string); ok {
		return s != ""
	}

	return true
}

func (i *Interpreter) isEqual(a, b any) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	switch av := a.(type) {
	case bool:
		bv, ok := b.(bool)
		return ok && av == bv

	case float64:
		bv, ok := b.(float64)
		return ok && av == bv

	case string:
		bv, ok := b.(string)
		return ok && av == bv
	}

	return a == b
}

func (i *Interpreter) checkNumberOperand(operator tokenizer.Token, operand any) {
	if _, ok := operand.(float64); ok {
		return
	}
	i.report(loxerror.NewRuntimeError(operator.Line(), operator.Lexeme(), loxerror.ErrExpectedNumericOperand))
}

func (i *Interpreter) checkNumberOperands(operator tokenizer.Token, left, right any) {
	_, lok := left.(float64)
	_, rok := right.(float64)

	if lok && rok {
		return
	}

	i.report(loxerror.NewRuntimeError(operator.Line(), operator.Lexeme(), loxerror.ErrExpectedNumericOperands))
}

func (i *Interpreter) checkMatchingOperandTypes(operator tokenizer.Token, left, right any) {
	var expectedType string
	switch left.(type) {
	case float64:
		if _, ok := right.(float64); ok {
			return
		}
		expectedType = "number"
		break
	case string:
		if _, ok := right.(string); ok {
			return
		}
		expectedType = "string"
		break
	}

	i.report(loxerror.NewRuntimeError(operator.Line(), fmt.Sprintf("%v", right), loxerror.ErrExpectedRightSideToBeOfLeftSideType(operator.Lexeme(), expectedType)))
}

func (i *Interpreter) stringify(value any) string {
	if value == nil {
		return "nil"
	}

	switch v := value.(type) {
	case bool:
		if v {
			return "true"
		}
		return "false"

	case float64:
		text := fmt.Sprintf("%g", v)

		if strings.ContainsAny(text, "eE") {
			return text
		}

		if strings.HasSuffix(text, ".0") {
			return text[:len(text)-2]
		}

		return text

	case string:
		return v
	}

	return fmt.Sprintf("%v", value)
}

func (i *Interpreter) report(err *loxerror.LoxError) {
	panic(err)
}

func (i *Interpreter) handleSuperSet(superExpr *ast.Super, nameTok tokenizer.Token, valueExpr ast.Expr) any {
	distance := i.locals[superExpr]
	superclass, ok := i.env.GetAt(distance, "super")
	if !ok {
		return nil
	}
	object, ok := i.env.GetAt(distance-1, "this")
	if !ok {
		return nil
	}
	value := i.evaluate(valueExpr)
	name := nameTok.Lexeme()

	// only search setters in *superclass chain*
	if setter, ok := superclass.(*Class).FindSetter(name); ok {
		return setter.Bind(object.(*ClassInstance)).Call(i, []any{value})
	}

	// assigning to a field via super is invalid
	i.report(loxerror.NewRuntimeError(nameTok.Line(), nameTok.Lexeme(), "cannot assign to superclass field"))
	return nil
}
