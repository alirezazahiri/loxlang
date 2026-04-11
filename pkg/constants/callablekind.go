package constants

type CallableKind string

const (
	CallableKindNone              CallableKind = "none"
	CallableKindFunction          CallableKind = "function"
	CallableKindInitializer       CallableKind = "initializer"
	CallableKindMethod            CallableKind = "method"
	CallableKindAnonymousFunction CallableKind = "anonymous"
)
