package main

import (
	"loxlang/internal/astgenerator"
	"loxlang/pkg/file"
)

func main() {
	outputDir := "internal/ast"

	expr := astgenerator.NewExpr()
	exprFileName, exprContent := expr.Generate()

	file.WriteFiles(outputDir, map[string][]byte{
		exprFileName: exprContent,
	})

}
