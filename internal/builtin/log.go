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

func (c *LogFn) Call(arguments []any) any {
	fmt.Println(arguments...)
	return nil
}

func (c *LogFn) Arity() int {
	return -1
}

func (c *LogFn) ToString() string {
	return "<native fn>"
}
