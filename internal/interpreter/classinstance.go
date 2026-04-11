package interpreter

import (
	"fmt"
	"loxlang/internal/environment"
	"loxlang/internal/tokenizer"
)

type ClassInstance struct {
	class  *Class
	fields map[string]any
}

func NewClassInstance(class *Class) *ClassInstance {
	return &ClassInstance{
		class:  class,
		fields: make(map[string]any),
	}
}

func (ci *ClassInstance) Get(closure *environment.Environment, property tokenizer.Token) (any, bool) {
	key := property.Lexeme()

	if v, ok := ci.fields[key]; ok {
		return v, true
	}

	if m, ok := ci.class.FindMethod(key); ok {
		return m.Bind(ci), true
	}
	
	return nil, false

}

func (ci *ClassInstance) Set(property tokenizer.Token, value any) {
	key := property.Lexeme()
	ci.fields[key] = value
}

func (ci *ClassInstance) String() string {
	return fmt.Sprintf("<instanceof %s>", ci.class.Name)
}
