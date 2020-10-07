# Overview

## Structs

- Useful for defining a unit with related information but different types
- Similar to C's structs
- `type <name> struct { ... }`

Eg.

```go
type Person struct {
  Name string
  Age int
}
```

- Struct-Literal syntax for initializing
  - `p := Person{"Gopher", 10}`
- Partial initialization?
- Fields can be accessed using the `.` notation
  - `p.Name` or `p.Age`
- `var p Person` zero value?
- Structs can be anonymous too!
  - `struct { ... }`
- Empty structs have 0 memory usage
  - `a := struct{}{}`
  - `a` has 0 memory foot-print
  - Useful for signaling

## Pointers to structs

- No dereferencing required
- Field access is simple for `p := &Person{}`
  - `p.Name = "Gopher"`
  vs.
  - `(*p).Name = "Gopher"`
