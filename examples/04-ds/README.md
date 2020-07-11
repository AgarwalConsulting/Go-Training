# Overview

## Arrays

- `var <name> [<length>]int`
- `len` & `cap` from `builtin`
- zero values?

## [Slices](https://blog.golang.org/slices-intro)

- `var <name> []int`
- References to underlying array
- `append`
  - Eg. `append(s, <element1>, <element2>, ...)`
  - You have to assign it back to `s` or whatever is your slice variable name
  - Can have a different behavior depending on capacity
- `make([]<type>, len, cap)`
- `make([]<type>, lenAndCap)`
- zero value?

## Maps

- `map` keyword
  - `map[<keyType>]<valType>`
- `make(map[string]Person)`
- `len` only
- `delete`
  - Eg. `delete(m, <key>)`
- lookup returns 2 values - `val, ok`
- zero value?
- They need to be initialized before usage either with `make` or `map` literal syntax: `map[string]Person{}`

## Custom

- generics are coming in Go 2!
