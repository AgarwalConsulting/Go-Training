# Overview

## Pass by `?`

- Go is pass by value always!
- A few types have pointer fields under the hood
  - Eg.
    - `slices`
    - `maps`
    - `channels`

## Pointers

- `p := &<variable>`
  - `var p *<type>`
- Print address using `%p` formatter
- `*` is used for dereferencing as well as defining a var as pointer
- `&` is used to get the address/reference of a variable
- `new` builtin function can initialize memory for any type

- No pointer arithmetic
  - You would have to use package [`unsafe`](https://golang.org/pkg/unsafe)
- `unsafe` is used heavily internally
