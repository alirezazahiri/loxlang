package interpreter

import (
	"fmt"
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

func (ci *ClassInstance) String() string {
	return fmt.Sprintf("<instanceof %s>", ci.class.Name)
}
