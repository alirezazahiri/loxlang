package builtin

import (
	"bufio"
	"loxlang/pkg/constants"
	"loxlang/pkg/interfaces"
	"os"
	"strings"
)

type InputFn struct {
	Name string
}

func NewInputFn() interfaces.LoxCallable {
	return &InputFn{
		Name: constants.INPUT_FN,
	}
}

func (i *InputFn) Call(arguments []any) any {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := strings.TrimSpace(scanner.Text())
	return line
}

func (i *InputFn) Arity() int {
	return 0
}

func (i *InputFn) ToString() string {
	return "<native fn>"
}
