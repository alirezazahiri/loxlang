package loxerror

import (
	"fmt"
)

type LoxError struct {
	Phase   Phase
	Line    int
	Where   string
	Message LoxErrorMessage
}

func (e *LoxError) Error() string {
	if e.Where != "" {
		return fmt.Sprintf("[%s] line %d at %q: %s", e.Phase, e.Line, e.Where, e.Message)
	}
	return fmt.Sprintf("[%s] line %d: %s", e.Phase, e.Line, e.Message)
}

func NewTokenizerError(line int, where string, message LoxErrorMessage) *LoxError {
	return &LoxError{
		Phase:   PhaseTokenizer,
		Line:    line,
		Where:   where,
		Message: message,
	}
}

func NewParserError(line int, where string, message LoxErrorMessage) *LoxError {
	return &LoxError{
		Phase:   PhaseParser,
		Line:    line,
		Where:   where,
		Message: message,
	}
}

func NewRuntimeError(line int, where string, message LoxErrorMessage) *LoxError {
	return &LoxError{
		Phase:   PhaseRuntime,
		Line:    line,
		Where:   where,
		Message: message,
	}
}
