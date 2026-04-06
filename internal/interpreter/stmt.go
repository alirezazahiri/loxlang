package interpreter

import (
	"fmt"
	"loxlang/internal/ast"
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
