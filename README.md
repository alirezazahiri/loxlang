# loxlang

A small dynamically typed language and interpreter written in Go. **loxlang** offers familiar imperative syntax, first-class functions, and object-oriented features including inheritance and property accessors.

If you have read [*Crafting Interpreters*](https://craftinginterpreters.com/), much of the core will feel familiar—loxlang is documented here on its own terms, with extensions such as `console`, `input`, and class getters/setters called out in the docs.

## Features

- Tree-walking interpreter: tokenizer, parser, resolver, and evaluator
- Lexical scoping, closures, and classes with single inheritance
- Property getters (`name { ... }`) and setters (`set name(v) { ... }`)
- Built-ins: `print`, `clock`, `input`, and global `console.log`
- Interactive REPL and single-file execution
- Example programs under [`examples/`](examples/)

## Quick start

**Requirements:** [Go](https://go.dev/) 1.26+

```bash
# Run a script
go run ./cmd examples/sample4.lox

# REPL (empty line or `exit` to quit)
go run ./cmd
```

With [Task](https://taskfile.dev/) installed:

```bash
task file FILE=examples/sample3.lox
task repl
task build    # outputs bin/lox.exe
```

## Documentation

| Guide | Description |
|-------|-------------|
| [Getting started](docs/getting-started.md) | Installation, CLI, REPL, errors, and examples |
| [Language guide](docs/language-guide.md) | Types, syntax, control flow, functions, and closures |
| [Classes](docs/classes.md) | Classes, `init`, inheritance, `super`, getters, and setters |
| [Built-ins](docs/builtins.md) | `print`, `clock`, `input`, and `console` |

## Project layout

```
cmd/          CLI entrypoint (run scripts or REPL)
internal/     Tokenizer, parser, resolver, interpreter, AST
examples/     Sample .lox programs
pkg/          Shared utilities and error types
```

### Developer commands

```bash
task test      # go test ./...
task fmt       # go fmt ./...
task lint      # go vet ./...
task print FILE=examples/sample1.lox   # print AST
```

## Examples

| File | Demonstrates |
|------|----------------|
| `examples/sample1.lox` | Control flow and early `return` |
| `examples/sample2.lox` | Inheritance, getters/setters, `super` |
| `examples/sample3.lox` | Closures and `console.log` |
| `examples/sample4.lox` | Recursion |

## License

This project is licensed under the [MIT License](LICENSE).
