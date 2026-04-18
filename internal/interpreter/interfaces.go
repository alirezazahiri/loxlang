package interpreter

type LoxCallable interface {
	Call(i *Interpreter, args []any) any
	Arity() int
}


type LoxBindable interface {
	LoxCallable
	Bind(instance *ClassInstance) LoxCallable
}