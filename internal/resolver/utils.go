package resolver

import (
	"loxlang/pkg/loxerror"
)

func (r *Resolver) report(err *loxerror.LoxError) {
	panic(err)
}

func (r *Resolver) existsInOuterScopes(name string) bool {
    scopes := r.scopes

    for i := len(scopes) - 2; i >= 0; i-- {
        if _, ok := scopes[i][name]; ok {
            return true
        }
    }
    return false
}
