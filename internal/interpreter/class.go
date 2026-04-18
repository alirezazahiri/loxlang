package interpreter

import (
	"fmt"
	"loxlang/pkg/constants"
)

type Class struct {
	Name    string
	Super   *Class
	Methods map[string]LoxBindable
	Getters map[string]*LoxFunction
	Setters map[string]*LoxFunction
}

func NewClass(name string, super *Class, methods map[string]LoxBindable, getters map[string]*LoxFunction, setters map[string]*LoxFunction) *Class {
	return &Class{
		Name:    name,
		Super:   super,
		Methods: methods,
		Getters: getters,
		Setters: setters,
	}
}

func (c *Class) FindMethod(name string) (LoxBindable, bool) {
	m, ok := c.Methods[name]
	if !ok && c.Super != nil {
		return c.Super.FindMethod(name)
	}
	if m == nil {
		return nil, ok
	}
	return m.(LoxBindable), ok
}

func (c *Class) FindGetter(name string) (*LoxFunction, bool) {
	key := constants.GETTER_PREFIX + name
	m, ok := c.Getters[key]
	if !ok && c.Super != nil {
		return c.Super.FindGetter(name)
	}
	return m, ok
}

func (c *Class) FindSetter(name string) (*LoxFunction, bool) {
	key := constants.SETTER_PREFIX + name
	m, ok := c.Setters[key]
	if !ok && c.Super != nil {
		return c.Super.FindSetter(name)
	}
	return m, ok
}

func (c *Class) Call(i *Interpreter, args []any) any {
	instance := NewClassInstance(c)

	if initializer, ok := c.FindMethod("init"); ok {
		return initializer.Bind(instance).Call(i, args)
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
