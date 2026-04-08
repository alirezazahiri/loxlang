package astgenerator

func NewStmt() *AstGenerator {
	return New("Stmt", []string{
		"If   			: Expr condition, Stmt thenBranch, Stmt elseBranch",
		"Block   		: []Stmt statements",
		"Expression   	: Expr expression",
		"Print   		: Expr expression",
		"Var        	: tokenizer.Token name, Expr initializer",
	}, "loxlang/internal/tokenizer")
}
