package astgenerator

import (
	"fmt"
	"go/format"
	"loxlang/pkg/utils"
	"strings"
)

// <typeName> : <fieldType> <fieldName>, <fieldType> <fieldName>, ...
func (a *AstGenerator) Generate() (string, []byte) {
	filename := fmt.Sprintf("%s.go", strings.ToLower(a.baseName))

	var b strings.Builder

	b.WriteString(a.buildHeader())

	visitors := make([]string, 0, len(a.types))

	for _, t := range a.types {
		structName, fields := parseType(t)

		structFields, funcParams, initializers := buildFields(fields)

		b.WriteString(buildStruct(structName, structFields))
		b.WriteString(buildConstructor(structName, funcParams, initializers))
		b.WriteString(buildAcceptMethod(structName, a.baseName))

		visitors = append(
			visitors,
			buildVisitorSignature(structName, a.baseName),
		)
	}

	b.WriteString(buildVisitorInterface(a.baseName, visitors))

	formatted, err := format.Source([]byte(b.String()))
	if err != nil {
		panic(err)
	}

	return filename, formatted
}

func (a *AstGenerator) buildHeader() string {
	var b strings.Builder

	b.WriteString("// THIS FILE IS AUTO GENERATED AND IS READ-ONLY (DO NOT MODIFY IT)\n\n")

	b.WriteString("package ast\n\n")

	if len(a.imports) > 0 {
		b.WriteString("import (\n")

		for _, imp := range a.imports {
			b.WriteString(fmt.Sprintf("\"%s\"\n", imp))
		}

		b.WriteString(")\n\n")
	}

	b.WriteString(fmt.Sprintf(
		"type %s interface {\nAccept(visitor %sVisitor) any\n}\n\n",
		a.baseName,
		a.baseName,
	))

	return b.String()
}

func parseType(t string) (string, string) {
	parts := strings.Split(t, ":")
	return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
}

func buildFields(fieldString string) (structFields, funcParams, initializers []string) {
	fields := strings.Split(fieldString, ", ")

	structFields = make([]string, 0, len(fields))
	funcParams = make([]string, 0, len(fields))
	initializers = make([]string, 0, len(fields))

	for _, field := range fields {
		parts := strings.Split(field, " ")

		fieldType := parts[0]
		fieldName := parts[1]
		capitalized := utils.CapFirst(fieldName)

		funcParams = append(funcParams,
			fmt.Sprintf("%s %s", fieldName, fieldType),
		)

		structFields = append(structFields,
			fmt.Sprintf("%s %s", capitalized, fieldType),
		)

		initializers = append(initializers,
			fmt.Sprintf("%s:\t%s", capitalized, fieldName),
		)
	}

	return
}

func buildStruct(name string, fields []string) string {
	return fmt.Sprintf(
		"type %s struct {\n%s\n}\n\n",
		name,
		strings.Join(fields, "\n"),
	)
}

func buildConstructor(name string, params, inits []string) string {
	return fmt.Sprintf(
		"func New%s(%s) *%s {\nreturn &%s{\n%s,\n}\n}\n\n",
		name,
		strings.Join(params, ", "),
		name,
		name,
		strings.Join(inits, ",\n"),
	)
}

func buildAcceptMethod(structName, baseName string) string {
	return fmt.Sprintf(
		"func (e *%s) Accept(visitor %sVisitor) any {\nreturn visitor.Visit%s%s(e)\n}\n\n",
		structName,
		baseName,
		structName,
		baseName,
	)
}

func buildVisitorSignature(structName, baseName string) string {
	lower := strings.ToLower(baseName)

	return fmt.Sprintf(
		"Visit%s%s(%s *%s) any",
		structName,
		baseName,
		lower,
		structName,
	)
}

func buildVisitorInterface(baseName string, visitors []string) string {
	return fmt.Sprintf(
		"type %sVisitor interface {\n%s\n}\n\n",
		baseName,
		strings.Join(visitors, "\n"),
	)
}
