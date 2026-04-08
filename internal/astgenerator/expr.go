package astgenerator

func NewExpr() *AstGenerator {
	return New("Expr", []string{
		"Assign		: tokenizer.Token name, Expr value",
		"Binary		: Expr left, tokenizer.Token operator, Expr right",
		"Logical	: Expr left, tokenizer.Token operator, Expr right",
		"Grouping	: Expr expression",
		"Literal	: any value",
		"Unary		: tokenizer.Token operator, Expr right",
		"Variable	: tokenizer.Token name",
	}, "loxlang/internal/tokenizer")
}
