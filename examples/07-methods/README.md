# Overview

Go doesn't support function overloading!

## Methods

- Methods are known as receivers or receiver functions in Go
- Receivers can be defined on any user-defined type in Go
- You can define multiple receivers with same name but on different receiver types, within a package
- They need to be defined in the same package as the `type` is defined in
- Getters & Setters are anti-patterns in Go

### Value receivers

- `func (<varName> <receiverType>) <funcNameA>() { ... }`
  - Eg. `func (f MyFloat) Abs() float64 { ... }`
- There is no `this` or `self` keyword in go, which could give you a hold of the variable on which you are dispatching a receiver.
  - They are instead passed in to the receiver, the same way you have arguments to a function

### Pointer receivers

- `func (<varName> *<receiverType>) <funcNameB>() { ... }`
- Pointer receivers can help in:
  - Avoiding copy of a large-ish struct variable
  - Update fields of a struct variable
- Same as fields, don't need any dereferencing to invoke

### Value & Pointer types

- The value & pointer receivers can be dispatched interchangeably:
  - `var v <receiverType>`
    - `v.funcNameA()` & `v.funcNameB()`
  - `p := &v`
    - `p.funcNameA()` & `p.funcNameB()`

- Eg.

  value receiver: `(v Vertex) Abs`

  pointer receiver: `(v *Vertex) Scale`

  value: `var vertex Vertex`

  pointer: `var p *Vertex`

```golang
  vertex.Abs()
  vertex.Scale()

  p.Abs()
  p.Scale()
```
