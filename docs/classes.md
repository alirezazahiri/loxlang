# Classes

loxlang provides object-oriented programming with classes, instances, single inheritance, and JavaScript-style property accessors.

## Defining a class

```lox
class Person {
  init(name) {
    this.name = name;
  }

  greet() {
    return "Hello, " + this.name;
  }
}
```

- **`init`** is the initializer (constructor). It runs when the class is **called** like a function.
- **`this`** refers to the instance being initialized or the receiver of a method call.

```lox
var p = Person("Ada");
print p.greet();
```

## Methods

Methods are functions declared inside the class body:

```lox
class Counter {
  init() {
    this.n = 0;
  }

  bump() {
    this.n = this.n + 1;
  }
}
```

## Inheritance

Subclass with `<`:

```lox
class Animal {
  speak() {
    return "...";
  }
}

class Dog < Animal {
  speak() {
    return "woof";
  }
}
```

## `super`

Inside a subclass, `super.method` invokes the superclass implementation of `method`:

```lox
class B < A {
  foo() {
    super.foo();
    print "from B";
  }
}
```

`super` is only valid for qualified method access (`super.name`), not as a standalone expression.

## Property getters and setters

loxlang extends plain method syntax with **accessors**. They are declared without the `fun` keyword.

### Getter

A getter is an identifier followed directly by a block (no parameter list):

```lox
class Box {
  init(v) { this.val = v; }

  value {
    return this.val;
  }
}

var b = Box(10);
print b.value;   // invokes getter
```

### Setter

A setter uses the `set` keyword and exactly one parameter:

```lox
class Box {
  init(v) { this.val = v; }

  value {
    return this.val;
  }

  set value(v) {
    this.val = v;
  }
}

var b = Box(1);
b.value = 42;
print b.value;
```

### Assignment and inheritance

- Reading `obj.prop` calls the getter if one exists; otherwise it reads a field.
- Assigning `obj.prop = x` calls the setter when defined.
- Subclasses may override getters/setters and use `super.prop` inside them for the superclass behavior (see `examples/sample2.lox`).

## Calling and arity

- Construct with `ClassName(arg1, arg2)`; arguments are passed to `init` if present.
- If there is no `init`, constructing with zero arguments is allowed; extra arguments are an error.
- Maximum arity for user functions and calls is **255** arguments (implementation limit).

## Related topics

- [Language guide](language-guide.md) — functions and `this` in broader context
- [Built-ins](builtins.md) — native classes such as `console`
