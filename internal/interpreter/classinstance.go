package interpreter

import (
	"fmt"
	"loxlang/internal/tokenizer"
)

type ClassInstance struct {
	class  *Class
	fields map[string]any
}

func NewClassInstance(class *Class) *ClassInstance {
	return &ClassInstance{
		class: class,
	}
}

func (ci *ClassInstance) Get(property tokenizer.Token) (any, bool) {
	key := property.Lexeme()
	v, ok := ci.fields[key]

	if !ok {
		m, ok := ci.class.FindMethod(key)
		return m, ok 
	}

	return v, true
}

func (ci *ClassInstance) Set(property tokenizer.Token, value any) {
	key := property.Lexeme()
	ci.fields[key] = value
}

func (ci *ClassInstance) String() string {
	return fmt.Sprintf("<instanceof %s>", ci.class.Name)
}
