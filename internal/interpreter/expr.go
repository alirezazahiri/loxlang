package interpreter

import (
	"loxlang/internal/ast"
	"loxlang/internal/environment"
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
	distance, ok := i.locals[expr]
	success := false

	if fn, ok := value.(*LoxFunction); ok {
		if fn.declaration.Name.Lexeme() == "" {
			name := expr.Name.Lexeme()

			fn.declaration.Name = tokenizer.NewToken(tokenizer.IDENTIFIER, name, nil, expr.Name.Line())

			inner := environment.New(fn.closure)
			inner.Define(name, fn)
			fn.closure = inner
		}
	}
	
	if !ok {
		success = i.env.AssignAt(distance, expr.Name, value)
	} else {
		success = i.globals.Assign(expr.Name, value)
	}

	if !success {
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

func (i *Interpreter) VisitCallExpr(expr *ast.Call) any {
	callee := i.evaluate(expr.Callee)
	arguments := make([]any, 0)
	for _, arg := range expr.Arguments {
		arguments = append(arguments, i.evaluate(arg))
	}

	if function, ok := callee.(LoxCallable); ok {
		arity := function.Arity()
		argsLength := len(arguments)

		if arity != -1 && arity != argsLength {
			i.report(loxerror.NewRuntimeError(expr.Paren.Line(), expr.Paren.Lexeme(), loxerror.ErrArgsCountConflict(arity, argsLength)))
		}

		return function.Call(i, arguments)
	}

	i.report(loxerror.NewRuntimeError(expr.Paren.Line(), expr.Paren.Lexeme(), loxerror.ErrCanOnlyCallFunctionsOrClasses))
	return nil
}

func (i *Interpreter) VisitLambdaExpr(expr *ast.Lambda) any {
	closure := i.env

	return NewLoxFunction(
		ast.NewFunction(tokenizer.NewToken(tokenizer.IDENTIFIER, "", nil, -1), expr.Params, expr.Body, false, false),
		closure,
		false,
	)
}

func (i *Interpreter) VisitGetExpr(expr *ast.Get) any {
	object := i.evaluate(expr.Object)
	instance, ok := object.(*ClassInstance)
	if !ok {
		i.report(loxerror.NewRuntimeError(expr.Name.Line(), expr.Name.Lexeme(), loxerror.ErrOnlyInstancesHaveProps))
		return nil
	}

	name := expr.Name.Lexeme()

	if getter, ok := instance.class.FindGetter(name); ok {
		return getter.Bind(instance).Call(i, nil)
	}

	if value, ok := instance.fields[name]; ok {
		return value
	}

	if method, ok := instance.class.FindMethod(name); ok {
		return method.Bind(instance)
	}

	i.report(loxerror.NewRuntimeError(expr.Name.Line(), expr.Name.Lexeme(), loxerror.ErrUndefinedProperty(name)))
	return nil
}

func (i *Interpreter) VisitSetExpr(expr *ast.Set) any {
	if superExpr, ok := expr.Object.(*ast.Super); ok {
		return i.handleSuperSet(superExpr, expr.Name, expr.Value)
	}

	object := i.evaluate(expr.Object)
	instance, ok := object.(*ClassInstance)
	if !ok {
		i.report(loxerror.NewRuntimeError(expr.Name.Line(), expr.Name.Lexeme(), loxerror.ErrOnlyInstancesHaveProps))
		return nil
	}

	value := i.evaluate(expr.Value)
	name := expr.Name.Lexeme()

	if setter, ok := instance.class.FindSetter(name); ok {
		return setter.Bind(instance).Call(i, []any{value})
	}

	instance.fields[name] = value
	return value
}

func (i *Interpreter) VisitThisExpr(expr *ast.This) any {
	v, ok := i.lookUpVariable(expr.Keyword, expr)
	if !ok {
		// ! this never happens, because we catch it in resolution steps
		return nil
	}
	return v
}

func (i *Interpreter) VisitSuperExpr(expr *ast.Super) any {
	var superClass *Class
	var object *ClassInstance

	if distance, ok := i.locals[expr]; ok {
		if v, ok := i.env.GetAt(distance, tokenizer.KeywordLexemes[tokenizer.SUPER]); ok {
			superClass = v.(*Class)
		}
		if v, ok := i.env.GetAt(distance-1, tokenizer.KeywordLexemes[tokenizer.THIS]); ok {
			object = v.(*ClassInstance)
		}
	}

	// fmt.Printf("\n----------------------------------\nsearching in super class for %s:\nMethods: %v,\nGetters: %v,\nSetters: %v\n----------------------------------\n", expr.Method.Lexeme(), superClass.Methods, superClass.Getters, superClass.Setters)
	if method, ok := superClass.FindMethod(expr.Method.Lexeme()); ok {
		// fmt.Println("found method", method)
		return method.Bind(object)
	}
	// if setter, ok := superClass.FindSetter(expr.Method.Lexeme()); ok {
	// 	fmt.Println("found setter", setter)
	// 	return setter.Bind(object)
	// }
	if getter, ok := superClass.FindGetter(expr.Method.Lexeme()); ok {
		// fmt.Println("found getter", getter)
		return getter.Bind(object).Call(i, nil)
	}

	i.report(loxerror.NewRuntimeError(expr.Method.Line(), expr.Method.Lexeme(), loxerror.ErrUndefinedProperty(expr.Method.Lexeme())))

	return nil
}
