package astgenerator

func NewStmt() *AstGenerator {
	return New("Stmt", []string{
		"Expression   	: Expr expression",
		"Print   		: Expr expression",
		"Var        	: tokenizer.Token name, Expr initializer",
	}, "loxlang/internal/tokenizer")
}
