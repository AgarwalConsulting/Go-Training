# Overview

## Packages

- Every `.go` file starts with `package <name>`
- All `.go` files in a directory need to belong to the same package
- Every package needs to be in a directory of the same name
  - Except `main` package
- `func main` can only be defined in a `main` package

## Exports

- Works using the first character -> uppercase means exported

## Imports

- Import path is always relative to `$GOPATH/src`

## Working with third-party packages

- `go get`
  - `go get -u all`
- `go mod`
  - `go mod init`
  - `go mod tidy`
  - `go mod why`
  - `go mod graph`
  - `go mod download`
- `go list -m all`

## More Reading

- `go help importpath`
