# Overview

Contrary to popular belief, go has inheritance through composition!

## Inheritance

### Struct

- A struct type can compose another type

For eg.

```go
type Bar string

func (b Bar) bar() {}

type Foo struct {
  Bar
  // Bar Bar
}
```

- Here the custom type `Foo` "inherits" Bar's bar method

- If a struct composes a struct, it can also "inherit" Fields

- Multiple inheritance is possible

### Interfaces

- An interface can compose another interface
