package astgenerator

func NewExpr() *AstGenerator {
	return New("Expr", []string{
		"Binary   : Expr left, tokenizer.Token operator, Expr right",
		"Grouping : Expr expression",
		"Literal  : any value",
		"Unary    : tokenizer.Token operator, Expr right",
	}, "loxlang/pkg/types")
}
