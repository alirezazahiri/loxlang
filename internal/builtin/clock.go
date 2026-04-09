package builtin

import (
	"loxlang/pkg/constants"
	"loxlang/pkg/interfaces"
	"time"
)

type ClockFn struct {
	Name string
}

func NewClockFn() interfaces.LoxCallable {
	return &ClockFn{
		Name: constants.CLOCK_FN,
	}
}

func (c *ClockFn) Call(arguments []any) any {
	return time.Now().Unix()
}

func (c *ClockFn) Arity() int {
	return 0
}

func (c *ClockFn) String() string {
	return "<native fn>"
}
