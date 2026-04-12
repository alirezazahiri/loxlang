package environment

import (
	"loxlang/internal/tokenizer"
)

func (e *Environment) Define(name string, value any) {
	e.values[name] = value
}

func (e *Environment) Assign(name tokenizer.Token, value any) bool {
	_, ok := e.values[name.Lexeme()]
	if ok {
		e.values[name.Lexeme()] = value
		return true
	}

	if e.enclosing != nil {
		return e.enclosing.Assign(name, value)
	}

	return false
}

func (e *Environment) AssignAt(distance int, name tokenizer.Token, value any) bool {
	ancestor := e.Ancestor(distance)
	if ancestor != nil {
		ancestor.Assign(name, value)
		return true
	}
	return false
}

func (e *Environment) Get(name string) (any, bool) {
	value, ok := e.values[name]

	if ok {
		return value, true
	}

	if e.enclosing != nil {
		return e.enclosing.Get(name)
	}

	return nil, false
}

func (e *Environment) GetAt(distance int, name string) (any, bool) {
	ancestor := e.Ancestor(distance)
	v, ok := ancestor.values[name]
	return v, ok
}

func (e *Environment) Ancestor(distance int) *Environment {
	env := e
	for range distance {
		env = env.enclosing
	}
	return env
}
