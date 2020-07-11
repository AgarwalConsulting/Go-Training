# Overview

- A poor version of exceptions

## Panic

- from `builtin` package, `panic` function
- Halts execution of the program
- Many things cause a panic
  - Eg.
    - Divide by zero
    - index out of bounds
    - nil pointer dereference
- Don't `recover` from panics, let it crash!

## Recover

A deferred `recover` fn can handle a potential panic.
