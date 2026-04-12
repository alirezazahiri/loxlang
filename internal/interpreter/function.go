package interpreter

import (
	"fmt"
	"loxlang/internal/ast"
	"loxlang/internal/environment"
	"loxlang/internal/tokenizer"
)

type LoxFunction struct {
	declaration   *ast.Function
	closure       *environment.Environment
	isInitializer bool
}

func NewLoxFunction(
	decl *ast.Function,
	closure *environment.Environment,
	isInitializer bool,
) *LoxFunction {
	return &LoxFunction{
		declaration:   decl,
		closure:       closure,
		isInitializer: isInitializer,
	}
}

func (f *LoxFunction) Bind(instance *ClassInstance) LoxCallable {
	env := environment.New(f.closure)
	env.Define(tokenizer.KeywordLexemes[tokenizer.THIS], instance)
	return NewLoxFunction(f.declaration, env, f.isInitializer)
}

func (f *LoxFunction) Call(i *Interpreter, args []any) (res any) {
	env := environment.New(f.closure)
	for idx, param := range f.declaration.Params {
		env.Define(param.Lexeme(), args[idx])
	}

	defer func() {
		if r := recover(); r != nil {
			if ret, ok := r.(*Return); ok {
				value := ret.Value
				if f.isInitializer {
					v, _ := f.closure.GetAt(0, tokenizer.KeywordLexemes[tokenizer.THIS])
					value = v
				}
				res = value
				return
			}
			panic(r)
		}
	}()

	i.executeBlock(f.declaration.Body, env)

	if f.isInitializer {
		v, _ := f.closure.GetAt(0, tokenizer.KeywordLexemes[tokenizer.THIS])
		return v
	}

	return nil
}

func (f *LoxFunction) Arity() int {
	return len(f.declaration.Params)
}

func (f *LoxFunction) String() string {
	if f.declaration.Name.Lexeme() != "" {
		return fmt.Sprintf("<fn %s>", f.declaration.Name.Lexeme())
	}
	return "<fn anonymous>"
}
