package constants

type CallableKind string

const (
	CallableKindNone              CallableKind = "none"
	CallableKindFunction          CallableKind = "function"
	CallableKindAnonymousFunction CallableKind = "anonymous"
)
