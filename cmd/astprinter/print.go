package main

import (
	"bufio"
	"fmt"
	"loxlang/internal/astprinter"
	"loxlang/internal/lox"
	"loxlang/pkg/file"
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
			panic("Usage: print [script]")
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
	run(l, string(bytes))
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

		run(l, line)
	}
}

func run(l *lox.LoxLang, content string) {
	tokens := l.Tokenize(content)

	if l.HasErrors() {
		fmt.Println("Tokenization errored!")
		return
	}

	ast := l.Parse(tokens)

	if l.HasErrors() {
		fmt.Println("Parsing errored!")
		return
	}

	printer := astprinter.New()

	fmt.Println(printer.Print(ast))
}
