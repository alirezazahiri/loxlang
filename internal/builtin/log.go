package builtin

import (
	"fmt"
	"loxlang/pkg/constants"
	"loxlang/pkg/interfaces"
)

type LogFn struct {
	Name string
}

func NewLogFn() interfaces.LoxCallable {
	return &LogFn{
		Name: constants.LOG_FN,
	}
}

func (l *LogFn) Call(arguments []any) any {
	fmt.Print(arguments...)
	return nil
}

func (l *LogFn) Arity() int {
	return -1
}

func (l *LogFn) ToString() string {
	return "<native fn>"
}
