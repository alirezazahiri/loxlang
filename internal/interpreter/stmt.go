package interpreter

import (
	"fmt"
	"loxlang/internal/ast"
	"loxlang/internal/environment"
)

func (i *Interpreter) VisitExpressionStmt(stmt *ast.Expression) any {
	i.evaluate(stmt.Expression)
	return nil
}
func (i *Interpreter) VisitPrintStmt(stmt *ast.Print) any {
	result := i.evaluate(stmt.Expression)
	fmt.Println(i.stringify(result))
	return nil
}
func (i *Interpreter) VisitVarStmt(stmt *ast.Var) any {
	var result any

	if stmt.Initializer != nil {
		result = i.evaluate(stmt.Initializer)
	}

	i.env.Define(stmt.Name.Lexeme(), result)
	return nil
}

func (i *Interpreter) VisitBlockStmt(stmt *ast.Block) any {
	i.executeBlock(stmt.Statements, environment.New(i.env))
	return nil
}

func (i *Interpreter) VisitIfStmt(stmt *ast.If) any {
	if i.isTruthy(i.evaluate(stmt.Condition)) {
		i.execute(stmt.ThenBranch)
	} else if stmt.ElseBranch != nil {
		i.execute(stmt.ElseBranch)
	}
	return nil
}

func (i *Interpreter) VisitWhileStmt(stmt *ast.While) any {
	cond := stmt.Condition
	for i.isTruthy(i.evaluate(cond)) {
		i.execute(stmt.Then)
	}
	return nil
}

func (i *Interpreter) VisitForStmt(stmt *ast.For) any {
	previous := i.env
	i.env = environment.New(previous)
	defer func() { i.env = previous }()

	init := stmt.Initializer
	cond := stmt.Condition
	mod := stmt.Modifier
	body := stmt.Body

	if init != nil {
		i.execute(init)
	}

	for {
		if cond != nil && !i.isTruthy(i.evaluate(cond)) {
			break
		}
		i.execute(body)
		if mod != nil {
			i.evaluate(mod)
		}
	}

	return nil
}

func (i *Interpreter) VisitFunctionStmt(stmt *ast.Function) any {
	closure := i.env

	fn := NewLoxFunction(
		stmt,
		closure,
		false,
	)

	i.env.Define(stmt.Name.Lexeme(), fn)

	return fn
}

func (i *Interpreter) VisitReturnStmt(stmt *ast.Return) any {
	val := any(nil)
	if stmt.Value != nil {
		val = i.evaluate(stmt.Value)
	}
	NewReturn(val).Throw()
	return nil
}

func (i *Interpreter) VisitClassStmt(stmt *ast.Class) any {
	i.env.Define(stmt.Name.Lexeme(), nil)

	closure := i.env
	
	methods := make(map[string]*LoxFunction, 0)
	getters := make(map[string]*LoxFunction, 0)
	setters := make(map[string]*LoxFunction, 0)

	for _, m := range stmt.Methods {
		name := m.Name.Lexeme()
		if m.IsGetter {
			getters[name] = NewLoxFunction(&m, closure, false)
		} else if m.IsSetter {
			setters[name] = NewLoxFunction(&m, closure, false)
		} else {
			methods[name] = NewLoxFunction(&m, closure, name == "init")
		}
	}

	class := NewClass(stmt.Name.Lexeme(), methods, getters, setters)
	i.env.Assign(stmt.Name, class)

	return nil
}
