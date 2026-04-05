package tokenizer

type TokenType uint8

const (
	// (
	LEFT_PAREN TokenType = iota + 1
	// )
	RIGHT_PAREN
	// {
	LEFT_BRACE
	// }
	RIGHT_BRACE
	// ,
	COMMA
	// .
	DOT
	// -
	MINUS
	// +
	PLUS
	// ;
	SEMICOLON
	// /
	SLASH
	// *
	STAR

	// !
	BANG
	// !=
	BANG_EQUAL
	// =
	EQUAL
	// ==
	EQUAL_EQUAL
	// >
	GREATER
	// >=
	GREATER_EQUAL
	// <
	LESS
	// <=
	LESS_EQUAL

	// identifier (e.g. variable names, etc.)
	IDENTIFIER

	// string literal
	STRING

	// number literal
	NUMBER

	// and
	AND
	// class
	CLASS
	// else
	ELSE
	// false
	FALSE
	// fun
	FUN
	// for
	FOR
	// if
	IF
	// nil
	NIL
	// or
	OR
	// print
	PRINT
	// return
	RETURN
	// super
	SUPER
	// this
	THIS
	// true
	TRUE
	// var
	VAR
	// while
	WHILE
	// end of file
	EOF
)
