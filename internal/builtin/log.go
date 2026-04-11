package builtin

import (
	"fmt"
	"loxlang/internal/interpreter"
	"loxlang/pkg/constants"
)

type LogFn struct {
	Name string
}

func NewLogFn() interpreter.LoxCallable {
	return &LogFn{
		Name: constants.LOG_FN,
	}
}

func (l *LogFn) Call(_ *interpreter.Interpreter, arguments []any) any {
	fmt.Println(arguments...)
	return nil
}

func (l *LogFn) Arity() int {
	return -1
}

func (l *LogFn) String() string {
	return "<native fn>"
}
