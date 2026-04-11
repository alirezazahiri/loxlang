package loxerror

import (
	"fmt"
	"loxlang/pkg/constants"
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
	ErrExpectedNumericOperand                       LoxErrorMessage = "expected numeric operand after operator"
	ErrExpectedNumericOperands                      LoxErrorMessage = "both sides of the operator must be numeric"
	ErrExpectedSemicolonAfterExpression             LoxErrorMessage = "expected \";\" after an expression"
	ErrExpectedIdenAfterVar                         LoxErrorMessage = "expected an identifier after using the \"var\" keyword"
	ErrExpectedInitializationOrSemicolon            LoxErrorMessage = "expected \";\" or initialization expression"
	ErrUndefinedVariable                            LoxErrorMessage = "variable is not defined"
	ErrUnexpectedAssignmentExpression               LoxErrorMessage = "invalid assignment expression"
	ErrCannotAssignToNonDeclaredVariables           LoxErrorMessage = "cannot assign to variables before declaration"
	ErrUnexpectedUnclosedBlock                      LoxErrorMessage = "expected \"}\" after the statement"
	ErrExpectedOpeningParenAfterIf                  LoxErrorMessage = "expected \"(\" after \"if\""
	ErrExpectedClosingParenAfterExpression          LoxErrorMessage = "expected \")\" after expression"
	ErrExpectedOpeningParenAfterWhile               LoxErrorMessage = "expected \"(\" after \"while\""
	ErrExpectedOpeningParenAfterFor                 LoxErrorMessage = "expected \"(\" after \"for\""
	ErrCanOnlyCallFunctionsOrClasses                LoxErrorMessage = "can only call functions or classes"
	ErrExpectedIdenAfterFun                         LoxErrorMessage = "expected identifier after  \"fun\" keyword"
	ErrExpectedIdenAfterSet                         LoxErrorMessage = "expected identifier after  \"set\" keyword"
	ErrExpectedSetKeywordBeforeSet					LoxErrorMessage = "expected \"set\" keyword"
	ErrExpectedIdenForClassMethods                  LoxErrorMessage = "class methods require an identifier"
	ErrExpectedOpeningBraceBeforeClassBody          LoxErrorMessage = "expected \"{\" before class body"
	ErrExpectedClosingBraceAfterClassBody           LoxErrorMessage = "expected \"}\" after class body"
	ErrExpectedClosingParenAfterParams              LoxErrorMessage = "expected \")\" after parameters"
	ErrExpectParameterName                          LoxErrorMessage = "expected name for parameter"
	ErrExprectedOpeningParenAfterAnonymousFunc      LoxErrorMessage = "expected \"(\" after anonymous function declaration (hint: add after \"fun\" keyword)"
	ErrExprectedClosingBraceBeforeAnonymousFuncBody LoxErrorMessage = "expected \"}\" before anonymous function body"
	ErrUnexpectedReturnFromNonFunction              LoxErrorMessage = "cannot return outside of a function"
	ErrUnexpectedInitReturn                         LoxErrorMessage = "cannot return a value from an initializer"
	ErrExpectedPropertyNameAfterDot                 LoxErrorMessage = "expected property after \".\""
	ErrOnlyInstancesHaveProps                       LoxErrorMessage = "only instances of classes have properties"
	ErrUnexpectedThisOutsideOfClass                 LoxErrorMessage = "cannot use 'this' outside of a class"
	ErrSettersExpectExactlyOneParam                 LoxErrorMessage = "setters expect exactly one parameter"
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

var ErrMaximumArgsCountExceeded LoxErrorMessage = LoxErrorMessage(fmt.Sprintf("maximum arguments count exceeded (max=%d)", constants.MAX_ARGUMENTS_COUNT))
var ErrMaximumParamsCountExceeded LoxErrorMessage = LoxErrorMessage(fmt.Sprintf("maximum parameters count exceeded (max=%d)", constants.MAX_ARGUMENTS_COUNT))

func ErrArgsCountConflict(expectedCount, receivedCount int) LoxErrorMessage {
	return LoxErrorMessage(fmt.Sprintf("expected %d argument(s), but got %d", expectedCount, receivedCount))
}

func ErrExpectedOpeningParenAfterCallableKindIden(kind constants.CallableKind) LoxErrorMessage {
	if kind == constants.CallableKindAnonymousFunction {
		return ErrExprectedOpeningParenAfterAnonymousFunc
	}
	return LoxErrorMessage(fmt.Sprintf("expected \"(\" after %s identifier", kind))
}

func ErrExpectedOpeningBraceBeforeCallableKindBody(kind constants.CallableKind) LoxErrorMessage {
	if kind == constants.CallableKindAnonymousFunction {
		return ErrExprectedClosingBraceBeforeAnonymousFuncBody
	}
	return LoxErrorMessage(fmt.Sprintf("expected \"{\" before %s body", kind))
}

func ErrCannotReadLocalVarInItsOwnInitializer(name string) LoxErrorMessage {
	return LoxErrorMessage(fmt.Sprintf("can't read local variable in its own initializer (reading %q)", name))
}

func ErrDuplicateVarDeclaration(name string) LoxErrorMessage {
	return LoxErrorMessage(fmt.Sprintf("variable with name %q already declared in this scope", name))
}

func ErrUndefinedProperty(name string) LoxErrorMessage {
	return LoxErrorMessage(fmt.Sprintf("undefined property %q", name))
}
