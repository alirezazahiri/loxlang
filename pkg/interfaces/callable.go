package interfaces

type LoxCallable interface {
	Call(args []any) any
	Arity() int
	String() string
}
