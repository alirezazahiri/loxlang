package resolver

import (
	"loxlang/internal/ast"
	"loxlang/pkg/loxerror"
)

func (r *Resolver) VisitAssignExpr(expr *ast.Assign) any {
	r.Resolve(expr.Value)
	r.resolveLocal(expr, expr.Name)
	return nil
}

func (r *Resolver) VisitBinaryExpr(expr *ast.Binary) any {
	r.Resolve(expr.Left)
	r.Resolve(expr.Right)
	return nil
}

func (r *Resolver) VisitLogicalExpr(expr *ast.Logical) any {
	r.Resolve(expr.Left)
	r.Resolve(expr.Right)
	return nil
}

func (r *Resolver) VisitGroupingExpr(expr *ast.Grouping) any {
	r.Resolve(expr.Expression)
	return nil
}

func (r *Resolver) VisitLiteralExpr(expr *ast.Literal) any {

	return nil
}

func (r *Resolver) VisitUnaryExpr(expr *ast.Unary) any {
	r.Resolve(expr.Right)
	return nil
}

func (r *Resolver) VisitLambdaExpr(expr *ast.Lambda) any { return nil }

func (r *Resolver) VisitCallExpr(expr *ast.Call) any {
	r.Resolve(expr.Callee)
	for _, arg := range expr.Arguments {
		r.Resolve(arg)
	}
	return nil
}

func (r *Resolver) VisitVariableExpr(expr *ast.Variable) any {
	if r.initializing != nil && expr.Name.Lexeme() == r.initializing.Lexeme() {
        if !r.existsInOuterScopes(expr.Name.Lexeme()) {
            r.report(loxerror.NewRuntimeError(
                expr.Name.Line(),
                expr.Name.Lexeme(),
                loxerror.ErrCannotReadLocalVarInItsOwnInitializer(expr.Name.Lexeme()),
            ))
        }
    }
	
	v, ok := r.scopes.Peek().Get(expr.Name.Lexeme())
	if !r.scopes.IsEmpty() &&
		ok && !v {
		r.report(loxerror.NewRuntimeError(expr.Name.Line(), expr.Name.Lexeme(), loxerror.ErrCannotReadLocalVarInItsOwnInitializer(expr.Name.Lexeme())))
	}
	r.resolveLocal(expr, expr.Name)
	return nil
}
