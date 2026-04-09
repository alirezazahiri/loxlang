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
	name := stmt.Name
	params := stmt.Params
	body := stmt.Body

	closure := i.env

	arity := len(params)

	i.RegisterCallable(
		NewCallable(
			name.Lexeme(),
			func(args []any) (res any) {
				fnEnv := environment.New(closure)

				for idx := range arity {
					fnEnv.Define(params[idx].Lexeme(), args[idx])
				}

				defer func() {
					if r := recover(); r != nil {
						if ret, ok := r.(*Return); ok {
							res = ret.Value
						} else {
							panic(r)
						}
					}
				}()

				i.executeBlock(body, fnEnv)
				return nil
			},
			arity,
		))

	return nil
}

func (i *Interpreter) VisitReturnStmt(stmt *ast.Return) any {
	val := any(nil)
	if stmt.Value != nil {
		val = i.evaluate(stmt.Value)
	}
	NewReturn(val).Throw()
	return nil
}
