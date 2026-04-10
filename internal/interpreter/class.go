package interpreter

import (
	"fmt"
	"loxlang/pkg/interfaces"
)

type Class struct {
	Name string
	Methods map[string]interfaces.LoxCallable
}

func NewClass(name string, methods map[string]interfaces.LoxCallable) *Class {
	return &Class{
		Name: name,
		Methods: methods,
	}
}

func (c *Class) FindMethod(key string) (interfaces.LoxCallable, bool) {
	m, ok := c.Methods[key]
	return m, ok
}

func (c *Class) Call(args []any) any {
	return NewClassInstance(c)
}

func (c *Class) Arity() int {
	return 0
}

func (c *Class) String() string {
	return fmt.Sprintf("<class %s>", c.Name)
}
