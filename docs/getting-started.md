# Getting started

This guide covers how to run **loxlang**, use the REPL, and interpret errors when something goes wrong.

## Installation

loxlang is run from source with Go—there is no separate publish step required for local use.

1. Clone the repository.
2. Ensure [Go](https://go.dev/dl/) **1.26 or newer** is on your `PATH`.
3. From the repository root, run scripts with `go run ./cmd <file>`.

To build a standalone binary:

```bash
task build
# or: go build -o bin/lox.exe ./cmd
bin/lox.exe examples/sample1.lox
```

## Running programs

Pass **exactly one** script path to the CLI:

```bash
go run ./cmd examples/sample4.lox
```

The interpreter reads the file, reports tokenizer/parser/runtime errors on stderr (via formatted error messages), and exits with a non-zero status when running a file that fails.

| Exit code | When |
|-----------|------|
| `65` | Tokenizer or parser error (while running a file) |
| `70` | Runtime error (while running a file) |

Success produces exit code `0` and any output from `print` or `console.log`.

## REPL

Start the REPL with no arguments:

```bash
go run ./cmd
```

You will see a `> ` prompt. Enter a statement or declaration (semicolons are required where the grammar expects them). To leave the REPL, enter a blank line or type `exit`.

The REPL evaluates one logical unit per line; it is useful for quick experiments, not for multi-line class bodies (use a file for larger programs).

## Example programs

The [`examples/`](../examples/) directory contains `.lox` files you can run as-is:

```bash
go run ./cmd examples/sample2.lox
go run ./cmd examples/sample3.lox
```

See the [root README](../README.md#examples) for a short description of each sample.

## Comments

Line comments begin with `//` and run to the end of the line:

```lox
var x = 1; // initializer
```

## Tooling (optional)

If you use [Task](https://taskfile.dev/), common workflows are defined in [`Taskfile.yml`](../Taskfile.yml):

| Task | Purpose |
|------|---------|
| `task file FILE=path.lox` | Run a script |
| `task repl` | Start the REPL |
| `task test` | Run Go tests |
| `task print FILE=path.lox` | Print the parsed AST (debugging) |

## Next steps

- [Language guide](language-guide.md) — syntax, types, and control flow
- [Classes](classes.md) — OOP features
- [Built-ins](builtins.md) — standard library surface
