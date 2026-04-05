package astgenerator

type AstGenerator struct {
	baseName string
	types    []string
	imports  []string
}

func New(baseName string, types []string, imports ...string) *AstGenerator {
	return &AstGenerator{
		baseName: baseName,
		types:    types,
		imports:  imports,
	}
}
