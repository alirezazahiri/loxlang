package loxerror

type Phase string

const (
	PhaseTokenizer Phase = "tokenizer"
	PhaseParser    Phase = "parser"
	PhaseRuntime   Phase = "runtime"
)
