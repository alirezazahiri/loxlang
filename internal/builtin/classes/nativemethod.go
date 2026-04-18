package builtinclasses

import (
	"fmt"
	"loxlang/internal/interpreter"
)

type NativeMethod struct {
	name  string
	arity int
	fn    func(instance *interpreter.ClassInstance, args []any) any
}

func NewNativeMethod(name string, arity int, fn func(instance *interpreter.ClassInstance, args []any) any) *NativeMethod {
	return &NativeMethod{name, arity, fn}
}

func (m *NativeMethod) Call(_ *interpreter.Interpreter, args []any) any {
	panic("NativeMethod.Call: should be called only after Bind")
}

func (m *NativeMethod) Bind(instance *interpreter.ClassInstance) interpreter.LoxCallable {
	return &NativeBoundMethod{
		method:   m,
		instance: instance,
	}
}

func (m *NativeMethod) Arity() int {
	return m.arity
}

func (m *NativeMethod) String() string {
	return fmt.Sprintf("<native method %s>", m.name)
}

type NativeBoundMethod struct {
	method   *NativeMethod
	instance *interpreter.ClassInstance
}

func (b *NativeBoundMethod) Call(_ *interpreter.Interpreter, args []any) any {
	return b.method.fn(b.instance, args)
}

func (b *NativeBoundMethod) Arity() int {
	return b.method.Arity()
}

func (b *NativeBoundMethod) String() string {
	return fmt.Sprintf("<native bound method %s>", b.method.name)
}
