package loxerror

import (
	"fmt"
)

type LoxErrorMessage string

// Sentinel errors for the tokenizer phase
const (
	// Tokenizer errors
	ErrUnterminatedString LoxErrorMessage = "unterminated string"
	ErrUnexpectedChar     LoxErrorMessage = "unexpected character"

	// Parser errors
	ErrUnterminatedParenthese         LoxErrorMessage = "unterminated parenthese"
	ErrUnexpectedTokenAfterExpression LoxErrorMessage = "unexpected token after expression"

	// Runtime errors
	ErrExpectedNumericOperand              LoxErrorMessage = "expected numeric operand after operator"
	ErrExpectedNumericOperands             LoxErrorMessage = "both sides of the operator must be numeric"
	ErrExpectedSemicolonAfterExpression    LoxErrorMessage = "expected \";\" after an expression"
	ErrExpectedIdenAfterVar                LoxErrorMessage = "expected an identifier after using the \"var\" keyword"
	ErrExpectedInitializationOrSemicolon   LoxErrorMessage = "expected \";\" or initialization expression"
	ErrUndefinedVariable                   LoxErrorMessage = "variable is not defined"
	ErrUnexpectedAssignmentExpression      LoxErrorMessage = "invalid assignment expression"
	ErrCannotAssignToNonDeclaredVariables  LoxErrorMessage = "cannot assign to variables before declaration"
	ErrUnexpectedUnclosedBlock             LoxErrorMessage = "expected \"}\" after the statement"
	ErrExpectedOpeningParenAfterIf         LoxErrorMessage = "expected \"(\" after \"if\""
	ErrExpectedClosingParenAfterExpression LoxErrorMessage = "expected \")\" after expression"
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
