package builtin

import (
	"bufio"
	"fmt"
	"loxlang/pkg/constants"
	"loxlang/pkg/interfaces"
	"os"
	"strconv"
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
	if len(arguments) > 0 {
		fmt.Print(arguments...)
	}
	scanner.Scan()
	line := scanner.Text()
	f, err := strconv.ParseFloat(line, 64)
	
	if err != nil {
		return line
	}

	return f
}

func (i *InputFn) Arity() int {
	return -1
}

func (i *InputFn) String() string {
	return "<native fn>"
}
