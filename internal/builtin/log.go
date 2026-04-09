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
	fmt.Println(arguments...)
	return nil
}

func (l *LogFn) Arity() int {
	return -1
}

func (l *LogFn) String() string {
	return "<native fn>"
}
