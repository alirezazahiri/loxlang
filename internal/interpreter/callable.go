package interpreter

import (
	"fmt"
	"loxlang/pkg/interfaces"
)

type Callable struct {
	Name  string
	cb    func(args []any) any
	arity int
}

func NewCallable(name string, cb func(args []any) any, arity int) interfaces.LoxCallable {
	return &Callable{
		Name:  name,
		cb:    cb,
		arity: arity,
	}
}

func (f *Callable) Call(arguments []any) any {
	return f.cb(arguments)
}

func (f *Callable) Arity() int {
	return f.arity
}

func (f *Callable) String() string {
	return fmt.Sprintf("<fn %s>", f.Name)
}

func (i *Interpreter) RegisterCallable(callable interfaces.LoxCallable) {
	fn := callable.(*Callable)
	i.env.Define(fn.Name, fn)
}
