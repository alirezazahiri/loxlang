package resolver

import (
	"loxlang/internal/ast"
	"loxlang/pkg/constants"
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

func (r *Resolver) VisitLambdaExpr(expr *ast.Lambda) any {
	enclosing := r.currentFunction
	r.currentFunction = constants.CallableKindAnonymousFunction
	r.beginScope()

	for _, p := range expr.Params {
		r.declare(p)
		r.define(p)
	}

	r.Resolve(expr.Body)

	r.endScope()
	r.currentFunction = enclosing
	return nil
}

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

	r.resolveLocal(expr, expr.Name)
	return nil
}

func (r *Resolver) VisitGetExpr(expr *ast.Get) any {
	r.Resolve(expr.Object)
	return nil
}

func (r *Resolver) VisitSetExpr(expr *ast.Set) any {
	r.Resolve(expr.Value)
	r.Resolve(expr.Object)
	return nil
}

func (r *Resolver) VisitThisExpr(expr *ast.This) any {
	if r.currentClass == constants.ClassKindNone {
		r.report(loxerror.NewRuntimeError(expr.Keyword.Line(), expr.Keyword.Lexeme(), loxerror.ErrUnexpectedThisOutsideOfClass))
	}
	r.resolveLocal(expr, expr.Keyword)
	return nil
}

func (r *Resolver) VisitSuperExpr(expr *ast.Super) any {
	if r.currentClass == constants.ClassKindNone {
		r.report(loxerror.NewRuntimeError(expr.Keyword.Line(), expr.Keyword.Lexeme(), loxerror.ErrUnexpectedSuperOutsideOfClass))
	}
	if r.currentClass != constants.ClassKindSubClass {
		r.report(loxerror.NewRuntimeError(expr.Keyword.Line(), expr.Keyword.Lexeme(), loxerror.ErrUnexpectedSuperOutsideOfSubClass))
	}
	r.resolveLocal(expr, expr.Keyword)
	return nil
}
