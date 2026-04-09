package interpreter

import (
	"fmt"
	"loxlang/pkg/interfaces"
)

type Callable struct {
	Name        string
	cb          func(args []any) any
	paramsCount int
}

func NewCallable(name string, cb func(args []any) any, paramsCount int) interfaces.LoxCallable {
	return &Callable{
		Name:        name,
		cb:          cb,
		paramsCount: paramsCount,
	}
}

func (f *Callable) Call(arguments []any) any {
	return f.cb(arguments)
}

func (f *Callable) Arity() int {
	return f.paramsCount
}

func (f *Callable) ToString() string {
	return fmt.Sprintf("%s fn", f.Name)
}

func (i *Interpreter) RegisterCallable(callable interfaces.LoxCallable) {
	fn := callable.(*Callable)
	i.env.Define(fn.Name, fn)
}
