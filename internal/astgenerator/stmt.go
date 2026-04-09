package astgenerator

func NewStmt() *AstGenerator {
	return New("Stmt", []string{
		"If   			: Expr condition, Stmt thenBranch, Stmt elseBranch",
		"For    		: Stmt initializer, Expr condition, Expr modifier, Stmt body",
		"While    		: Expr condition, Stmt then",
		"Block   		: []Stmt statements",
		"Function   	: tokenizer.Token name, []tokenizer.Token params, []Stmt body",
		"Expression   	: Expr expression",
		"Print   		: Expr expression",
		"Var        	: tokenizer.Token name, Expr initializer",
	}, "loxlang/internal/tokenizer")
}
