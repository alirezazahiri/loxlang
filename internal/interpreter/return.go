package interpreter

type Return struct {
	Value any
}

func NewReturn(value any) *Return {
	return &Return{
		Value: value,
	}
}

func (r *Return) Throw() {
	panic(r)
}
