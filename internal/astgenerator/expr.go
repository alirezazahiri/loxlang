package astgenerator

func NewExpr() *AstGenerator {
	return New("Expr", []string{
		"Assign		: tokenizer.Token name, Expr value",
		"Binary		: Expr left, tokenizer.Token operator, Expr right",
		"Logical	: Expr left, tokenizer.Token operator, Expr right",
		"Grouping	: Expr expression",
		"Literal	: any value",
		"Super		: tokenizer.Token keyword, tokenizer.Token method",
		"This		: tokenizer.Token keyword",
		"Unary		: tokenizer.Token operator, Expr right",
		"Lambda		: []tokenizer.Token params, []Stmt body",
		"Call		: Expr callee, tokenizer.Token paren, []Expr arguments",
		"Get		: Expr object, tokenizer.Token name",
		"Set		: Expr object, tokenizer.Token name, Expr value",
		"Variable	: tokenizer.Token name",
	}, "loxlang/internal/tokenizer")
}
