# Built-ins

Globals are registered before your program runs. This page lists the standard surface area beyond user-defined functions and classes.

## `print`

```lox
print expr;
```

Evaluates `expr`, converts the result to a string, and prints one line to **stdout**. Handy for quick output; for multiple values or logging, see `console` below.

## `clock`

```lox
var t0 = clock();
// ... work ...
var t1 = clock();
print t1 - t0;
```

`clock()` takes no arguments and returns the current Unix time in **seconds** as a number (integer-valued in practice).

## `input`

```lox
var line = input("Name: ");
```

- With **no** arguments, reads one line from **stdin** (no prompt).
- With one or more arguments, prints them (like a prompt) then reads a line.

If the line parses as a number, `input` returns a number; otherwise it returns the string as read (trimming is not applied beyond what the scanner provides).

Use this in file scripts or the REPL when you need interactive input.

## `console`

A global **instance** (not a class you construct) with a native logging method:

```lox
console.log("a", 1, true);
```

`console.log` accepts any number of arguments, prints them separated by spaces, and returns `nil`. It is similar to `print` but follows variadic console semantics and is used throughout the [`examples/`](../examples/) closure demo.

There is no `console` class visible in user code—only the pre-bound `console` object.

## Error reporting

Built-ins follow the same runtime rules as user code: wrong arity or invalid operands produce **runtime errors** with line information. When executing a file, the CLI exits with code **70** on uncaught runtime errors.

## Related topics

- [Getting started](getting-started.md) — stdin/stdout and exit codes
- [Language guide](language-guide.md) — calling and arity in general
