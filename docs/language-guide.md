# Language guide

loxlang is a small, dynamically typed language. Types exist only at runtime; variables are introduced with `var`, and values include numbers, strings, Booleans, `nil`, functions, and class instances.

## Literals and types

| Syntax | Type | Notes |
|--------|------|--------|
| `123`, `3.14` | Number | IEEE-754 doubles (Go `float64`) |
| `"hello"` | String | Double-quoted; `\n` is not special in the lexer—use actual newlines in source sparingly |
| `true`, `false` | Boolean | |
| `nil` | Nil | Absence of a value |

### Truthiness

Conditions treat these as **false**: `nil`, `false`, numeric `0`, and `""`. Every other value is **true**, including non-zero numbers, non-empty strings, functions, and objects.

### Equality

`==` and `!=` compare values with type-specific rules (numbers with numbers, strings with strings, etc.). Mixed-type equality is not used for unlike types in typical patterns—operands should match expectations.

## Operators

| Precedence (high → low) | Operators |
|-------------------------|-----------|
| Unary | `!`, `-` |
| Multiplicative | `*`, `/` |
| Additive | `+`, `-` |
| Comparison | `>`, `>=`, `<`, `<=` |
| Equality | `==`, `!=` |
| Logical AND | `and` |
| Logical OR | `or` |

`+` concatenates two strings or adds two numbers. Other arithmetic operators require numeric operands (runtime error otherwise).

## Variables and assignment

```lox
var a = 1;
var name;          // defaults to nil
a = a + 1;
```

Assignment targets must be a declared variable, object field, or property setter (see [Classes](classes.md)).

## Statements

### Expression statements

Any expression followed by `;`:

```lox
clock();
```

### Print

```lox
print "hello";
print 1 + 2;
```

`print` converts its operand to a readable string and writes one line to stdout.

### Blocks

```lox
{
  var x = 1;
  print x;
}
```

Blocks introduce a new scope.

### If / else

```lox
if (condition) {
  print "yes";
} else {
  print "no";
}
```

### While

```lox
var i = 0;
while (i < 3) {
  print i;
  i = i + 1;
}
```

### For

C-style three-part `for` loops are supported:

```lox
for (var i = 0; i < 3; i = i + 1) {
  print i;
}
```

Any of the initializer, condition, or increment may be omitted, following the usual C-like pattern (e.g. `for (;;)` with an internal `return` if you need to stop).

### Return

```lox
fun greet() {
  return "hi";
}
```

A bare `return;` returns `nil`. `return` with a value is illegal inside class `init` methods.

## Functions

### Declarations

```lox
fun add(a, b) {
  return a + b;
}

print add(1, 2);
```

Functions are values: a declared function is bound to its name in the current scope. Arity is fixed; calling with the wrong number of arguments is a runtime error.

### Anonymous functions (lambdas)

In expression position, `fun` starts a lambda:

```lox
var twice = fun(n) { return n * 2; };
print twice(21);
```

Lambdas close over their defining environment, so they support **closures**:

```lox
var make = fun(start) {
  return fun() {
    var prev = start;
    start = prev + 1;
    return prev;
  };
};

var next = make(1);
print next(); // 1
print next(); // 2
```

See `examples/sample3.lox` for a richer closure example.

### Calling

Call with parentheses: `callee(arg1, arg2)`. The callee must be a function or a class (classes construct instances—see [Classes](classes.md)).

## Member access

Use `.` for fields and methods on instances:

```lox
obj.field;
obj.method(arg);
```

## Related topics

- [Classes](classes.md) — methods, `this`, inheritance
- [Built-ins](builtins.md) — globals such as `clock` and `console`
