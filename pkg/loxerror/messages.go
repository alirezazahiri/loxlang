package loxerror

import (
	"fmt"
)

type LoxErrorMessage string

// Sentinel errors for the tokenizer phase
const (
	// Tokenizer errors
	ErrUnterminatedString LoxErrorMessage = "unterminated string"
	ErrUnexpectedChar                     = "unexpected character"

	// Parser errors
	ErrUnterminatedParenthese         = "unterminated parenthese"
	ErrUnexpectedTokenAfterExpression = "unexpected token after expression"

	// Runtime errors
	ErrExpectedNumericOperand  = "expected numeric operand after operator"
	ErrExpectedNumericOperands = "both sides of the operator must be numeric"
)

func ErrUnexpectedToken(received, expected string) LoxErrorMessage {
	if received == "" {
		return ErrUnterminatedParenthese
	}
	return LoxErrorMessage(fmt.Sprintf("unexpected token %q, expected: %q", received, expected))
}

func ErrExpectedRightSideToBeOfLeftSideType(operator, expectedType string) LoxErrorMessage {
	return LoxErrorMessage(fmt.Sprintf("expected the right side of the operator %q to be of type %q", operator, expectedType))
}
