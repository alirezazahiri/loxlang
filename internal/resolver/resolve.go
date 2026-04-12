package resolver

import (
	"loxlang/internal/ast"
	"loxlang/internal/scope"
	"loxlang/internal/tokenizer"
	"loxlang/pkg/constants"
	"loxlang/pkg/loxerror"
)

// Resolve panics if there is a resolution error
func (r *Resolver) Resolve(resolvable any) {
	switch resolvable.(type) {
	case []ast.Stmt:
		for _, statement := range resolvable.([]ast.Stmt) {
			r.Resolve(statement)
		}
		return
	case ast.Stmt:
		resolvable.(ast.Stmt).Accept(r)
		return
	case ast.Expr:
		resolvable.(ast.Expr).Accept(r)
	}
}

func (r *Resolver) resolveLocal(expr ast.Expr, name tokenizer.Token) {
	for i := r.scopes.Size() - 1; i >= 0; i-- {
		if r.scopes.Get(i).ContainsKey(name.Lexeme()) {
			r.interpreter.Resolve(expr, r.scopes.Size()-1-i)
			return
		}
	}
}

func (r *Resolver) resolveFunction(stmt ast.Function, kind constants.CallableKind) {
	enclosingFn := r.currentFunction
	r.currentFunction = kind

	r.beginScope()
	for _, param := range stmt.Params {
		r.declare(param)
		r.define(param)
	}
	r.Resolve(stmt.Body)
	r.endScope()

	r.currentFunction = enclosingFn
}

func (r *Resolver) resolveClass(stmt ast.Class) {
	r.beginScope()
	r.scopes.Peek().Put(tokenizer.KeywordLexemes[tokenizer.THIS], true)
	
	for _, fn := range stmt.Methods {
		kind := constants.CallableKindMethod
		if (fn.Name.Lexeme() == "init") {
			kind = constants.CallableKindInitializer
		} else if fn.IsGetter {
			kind = constants.CallableKindGetter
		} else if fn.IsSetter {
			kind = constants.CallableKindSetter
		}
		r.resolveFunction(fn, kind)
	}
	
	r.endScope()
}

func (r *Resolver) beginScope() {
	r.scopes = append(r.scopes, make(scope.Scope))
}

func (r *Resolver) endScope() {
	(&r.scopes).Pop()
}

func (r *Resolver) declare(name tokenizer.Token) {
	if r.scopes.IsEmpty() {
		return
	}
	scope := r.scopes.Peek()

	if scope.ContainsKey(name.Lexeme()) {
		r.report(loxerror.NewRuntimeError(name.Line(), name.Lexeme(), loxerror.ErrDuplicateVarDeclaration(name.Lexeme())))
	}

	scope.Put(name.Lexeme(), false)
}

func (r *Resolver) define(name tokenizer.Token) {
	if r.scopes.IsEmpty() {
		return
	}
	r.scopes.Peek().Put(name.Lexeme(), true)
}
