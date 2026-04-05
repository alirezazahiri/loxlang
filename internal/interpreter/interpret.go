package interpreter

import (
	"fmt"
	"loxlang/internal/ast"
)

func (i *Interpreter) Interpret(st ast.Expr) {
	result := i.evaluate(st)
	fmt.Println(i.stringify(result))
}
