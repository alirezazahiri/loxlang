package main

import (
	"bufio"
	"fmt"
	"loxlang/internal/lox"
	"loxlang/pkg/file"
	"loxlang/pkg/utils"
	"os"
	"strings"
)

func main() {
	args := os.Args

	files := make([]string, len(args))

	if len(args) > 1 {
		files = args[1:]
		if files != nil && len(files) == 1 {
			runFile(files[0])
		} else if len(files) > 1 {
			panic("Usage: loxlang [script]")
		}
	} else {
		runPrompt()
	}

	return
}

func runFile(path string) {
	f := file.New(path)

	bytes, fileReadError := f.Read()

	if fileReadError != nil {
		panic(fileReadError)
	}

	l := lox.New()
	fileContent := string(bytes)
	
	errs := l.Run(fileContent)
	utils.LogErrors(errs, true)
}

func runPrompt() {
	l := lox.New()
	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := strings.TrimSpace(scanner.Text())

		if line == "" || line == "exit" {
			return
		}

		errs := l.Run(line)
		utils.LogErrors(errs)
	}
}
