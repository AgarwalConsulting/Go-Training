# Overview

Interfaces are implicit in Go!

## Interface

- `interface` keyword
- `type Foo interface { ... }`
- It contains only method definitions.
- A type which implements all of the methods in an interface, implements the interface implicitly
- If any of the receiver methods, required by interface `I` on type `A` are pointer receiver, then only `*A` will be implementing an interface.
  - Eg.

```go
type I interface {
  Foo()
  Bar()
}

type A struct{}

func (a *A) Foo() {}

func (b *B) Bar() {}
```

then,

```go
var i Interface

// i = A{} // Will result in compiler error

i = &A{}
```

- Zero value of `i`?
- Interfaces are a combination of concrete type info and value
  - Eg.

```go
var a *A
var i interface = a

if i == nil {
  // Will this execute?
}
```

- There are tons of interfaces defined in the standard library. Some of them below:
  - `fmt.Stringer`
  - `io.Reader`
  - `io.Writer`
  - `json.Marshaler`
  - `error`

## Assertion

- With a `var` of interface type, you can print the underlying type info using `%T` formatter
- You can also get the underlying type variable using:
  - `i.(<underlying-type>)`
- You also have `switch v := i.(type) {`
  - where each case is for a specific type. Eg. `case int: `
