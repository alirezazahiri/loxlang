package builtinclasses

import "loxlang/internal/interpreter"


func Register(builder func() *interpreter.Class) *interpreter.ClassInstance {
	return interpreter.NewClassInstance(builder())
}