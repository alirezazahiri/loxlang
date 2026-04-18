package builtinclasses

import (
	"fmt"
	"loxlang/internal/interpreter"
	"loxlang/pkg/constants"
)

const (
	LOG_METHOD = "log"
)

func NewConsoleClass() *interpreter.Class {
	return interpreter.NewClass(
		constants.CONSOLE_CLASS,
		nil,
		map[string]interpreter.LoxBindable{
			LOG_METHOD: LogMethod(),
		},
		map[string]*interpreter.LoxFunction{},
		map[string]*interpreter.LoxFunction{},
	)
}

func LogMethod() *NativeMethod {
	return NewNativeMethod(
		LOG_METHOD,
		-1,
		func(_ *interpreter.ClassInstance, args []any) any {
			fmt.Println(args...)
			return nil
		},
	)
}
