package constants

type CallableKind string

const (
	CallableKindNone              CallableKind = "none"
	CallableKindFunction          CallableKind = "function"
	CallableKindMethod            CallableKind = "method"
	CallableKindAnonymousFunction CallableKind = "anonymous"
)
