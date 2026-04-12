package resolver

import (
	"loxlang/internal/ast"
	"loxlang/internal/tokenizer"
	"loxlang/pkg/constants"
	"loxlang/pkg/loxerror"
)

func (r *Resolver) VisitIfStmt(stmt *ast.If) any {
	r.Resolve(stmt.Condition)
	r.Resolve(stmt.ThenBranch)
	if stmt.ElseBranch != nil {
		r.Resolve(stmt.ElseBranch)
	}
	return nil
}

func (r *Resolver) VisitForStmt(stmt *ast.For) any {
	if stmt.Initializer != nil {
		r.Resolve(stmt.Initializer)
	}
	if stmt.Condition != nil {
		r.Resolve(stmt.Condition)
	}
	if stmt.Modifier != nil {
		r.Resolve(stmt.Modifier)
	}
	return nil
}

func (r *Resolver) VisitWhileStmt(stmt *ast.While) any {
	r.Resolve(stmt.Condition)
	r.Resolve(stmt.Then)
	return nil
}

func (r *Resolver) VisitBlockStmt(stmt *ast.Block) any {
	r.beginScope()
	r.Resolve(stmt.Statements)
	r.endScope()
	return nil
}

func (r *Resolver) VisitFunctionStmt(stmt *ast.Function) any {
	r.declare(stmt.Name)
	r.define(stmt.Name)

	r.resolveFunction(*stmt, constants.CallableKindFunction)
	return nil
}

func (r *Resolver) VisitClassStmt(stmt *ast.Class) any {
	enclosingClass := r.currentClass
	r.currentClass = constants.ClassKindClass

	r.declare(stmt.Name)
	r.define(stmt.Name)

	if stmt.SuperClass != nil && stmt.Name.Lexeme() == stmt.SuperClass.Name.Lexeme() {
		r.report(loxerror.NewRuntimeError(stmt.Name.Line(), stmt.SuperClass.Name.Lexeme(), loxerror.ErrClassesCannotInheritThemselves))
	}

	if stmt.SuperClass != nil {
		r.currentClass = constants.ClassKindSubClass
		r.Resolve(stmt.SuperClass)
		r.beginScope()
		r.scopes.Peek().Put(tokenizer.KeywordLexemes[tokenizer.SUPER], true)
	}

	r.resolveClass(*stmt)

	if stmt.SuperClass != nil {
		r.endScope()
	}

	r.currentClass = enclosingClass

	return nil
}

func (r *Resolver) VisitExpressionStmt(stmt *ast.Expression) any {
	r.Resolve(stmt.Expression)
	return nil
}

func (r *Resolver) VisitReturnStmt(stmt *ast.Return) any {
	if r.currentFunction == constants.CallableKindNone {
		r.report(loxerror.NewRuntimeError(stmt.Keyword.Line(), stmt.Keyword.Lexeme(), loxerror.ErrUnexpectedReturnFromNonFunction))
	}
	if stmt.Value != nil {
		if r.currentFunction == constants.CallableKindInitializer {
			r.report(loxerror.NewRuntimeError(stmt.Keyword.Line(), stmt.Keyword.Lexeme(), loxerror.ErrUnexpectedInitReturn))
		}
		r.Resolve(stmt.Value)
	}
	return nil
}

func (r *Resolver) VisitPrintStmt(stmt *ast.Print) any {
	r.Resolve(stmt.Expression)
	return nil
}

func (r *Resolver) VisitVarStmt(stmt *ast.Var) any {
	r.declare(stmt.Name)

	prev := r.initializing
	r.initializing = &stmt.Name

	if stmt.Initializer != nil {
		r.Resolve(stmt.Initializer)
	}

	r.initializing = prev

	r.define(stmt.Name)
	return nil
}
