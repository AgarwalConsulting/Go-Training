# Overview

Go is a statically typed, type inferred language!

i.e.) You don't have to write a type of a variable explicitly, the compile will infer it. But all types must be known/inferred at compile time.

## Variables

- `var` keyword

### Declaration & Initialization

- `var <name> <type>`
- `var <name> = <val>`
- `<name> := <val>`

### Default values & types

- zero values
- [Basic types](https://tour.golang.org/basics/11)
  - defined in the [`builtin`](https://golang.org/pkg/builtin/) package
- multiple assignment
  - `var i, j = "Hello", 42`

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

- Know type information using `%T` [formatter](https://golang.org/pkg/fmt/)
- Type aliases using `type` keyword

### Type conversion

- `<type>(<val-of-other-type>)`
  - Eg: `val := int(42.00)`
- Specialized type conversion `strconv`
