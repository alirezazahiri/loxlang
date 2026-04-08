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
