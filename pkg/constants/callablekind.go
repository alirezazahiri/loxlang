package constants

type CallableKind string

const (
	CallableKindNone              CallableKind = "none"
	CallableKindSetter            CallableKind = "setter"
	CallableKindGetter            CallableKind = "getter"
	CallableKindFunction          CallableKind = "function"
	CallableKindInitializer       CallableKind = "initializer"
	CallableKindMethod            CallableKind = "method"
	CallableKindAnonymousFunction CallableKind = "anonymous"
)
