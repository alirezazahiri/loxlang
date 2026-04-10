package interpreter

import (
	"fmt"
	"loxlang/internal/ast"
	"loxlang/internal/environment"
	"loxlang/internal/tokenizer"
	"loxlang/pkg/interfaces"
)

type CallableCallbackFn func(args []any) any

type Callable struct {
	Name  string
	cb    CallableCallbackFn
	arity int
}

func NewCallable(name string, cb CallableCallbackFn, arity int) interfaces.LoxCallable {
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
	if f.Name != "" {
		return fmt.Sprintf("<fn %s>", f.Name)
	}
	return "<fn anonymous>"
}

func (i *Interpreter) RegisterCallable(callable interfaces.LoxCallable) {
	fn := callable.(*Callable)
	i.env.Define(fn.Name, fn)
}

func (i *Interpreter) CallableCallback(closure *environment.Environment, params []tokenizer.Token, body []ast.Stmt) CallableCallbackFn {
	return func(args []any) (res any) {
		fnEnv := environment.New(closure)

		for idx := range len(params) {
			fnEnv.Define(params[idx].Lexeme(), args[idx])
		}

		defer func() {
			if r := recover(); r != nil {
				if ret, ok := r.(*Return); ok {
					res = ret.Value
				} else {
					panic(r)
				}
			}
		}()

		i.executeBlock(body, fnEnv)
		return nil
	}
}
