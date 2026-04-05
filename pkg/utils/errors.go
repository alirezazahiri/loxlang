package utils

import (
	"fmt"
	"loxlang/pkg/loxerror"
	"os"
)

func LogErrors(errs []*loxerror.LoxError, exit ...bool) {
	shouldExit := false
	
	if exit != nil {
		shouldExit = exit[0]
	}

	if errs != nil {
		for _, err := range errs {
			fmt.Println(err)
		}

		if errs[0].Phase == loxerror.PhaseRuntime && shouldExit {
			os.Exit(70)
		}

		if shouldExit {
			os.Exit(65)
		}
	}
}
