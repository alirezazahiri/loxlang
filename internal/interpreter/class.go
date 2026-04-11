package interpreter

import (
	"fmt"
)

type Class struct {
	Name    string
	Methods map[string]*LoxFunction
	Getters map[string]*LoxFunction
	Setters map[string]*LoxFunction
}

func NewClass(name string, methods map[string]*LoxFunction, getters map[string]*LoxFunction, setters map[string]*LoxFunction) *Class {
	return &Class{
		Name:    name,
		Methods: methods,
		Getters: getters,
		Setters: setters,
	}
}

func (c *Class) FindMethod(name string) (*LoxFunction, bool) {
	m, ok := c.Methods[name]
	return m, ok
}

func (c *Class) FindGetter(name string) (*LoxFunction, bool) {
	m, ok := c.Getters[name]
	return m, ok
}

func (c *Class) FindSetter(name string) (*LoxFunction, bool) {
	m, ok := c.Setters[name]
	return m, ok
}

func (c *Class) Call(i *Interpreter, args []any) any {
	instance := NewClassInstance(c)

	if initializer, ok := c.FindMethod("init"); ok {
		initializer.Bind(instance).Call(i, args)
	}

	return instance
}

func (c *Class) Arity() int {
	if initializer, ok := c.FindMethod("init"); ok {
		return initializer.Arity()
	}
	return 0
}

func (c *Class) String() string {
	return fmt.Sprintf("<class %s>", c.Name)
}
