package builtin

import (
	"loxlang/internal/interpreter"
	"loxlang/pkg/constants"
	"time"
)

type ClockFn struct {
	Name string
}

func NewClockFn() interpreter.LoxCallable {
	return &ClockFn{
		Name: constants.CLOCK_FN,
	}
}

func (c *ClockFn) Call(_ *interpreter.Interpreter, arguments []any) any {
	return time.Now().Unix()
}
 
func (c *ClockFn) Arity() int {
	return 0
}

func (c *ClockFn) String() string {
	return "<native fn>"
}
