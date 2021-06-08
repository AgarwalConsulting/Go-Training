# Overview

Errors are values in Go!

## Errors

- Most of the funcs in the standard library return 2 values. One is the expected value of expected type, the other is `error`.
- `error` is an interface defined in the `builtin` package
  - `type error interface { Error() string }`
  - Any type which has a `Error() string` receiver, implements it
- Standard library has some concrete implementations of `error` interface. Like:
  - `*fs.PathError`
  - `*strconv.NumError`
- package [`errors`](https://golang.org/pkg/errors/)
  - `errors.New(<message>)`
    - it makes it easy to create errors on the fly

## Russian doll of errors

- With Go 1.13, you can define custom errors which wrap other errors like a Matryoshka...

![Matryoshka](assets/matryoshka.jpg)

- `errors` package has 2 new methods:
  - `As` finds the first error in err's chain that matches target
  - `Is` reports whether any error in err's chain matches target

- All you need to do, in order for standard library to work with such an implementation of `error` is to implement `Unwrap() error` method on your custom error type
