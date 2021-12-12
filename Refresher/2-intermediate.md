# Intermediate - Refresher

- `panic` & `recover`

- Functional Programming in Go
  - Functions `func (<argName> <type>) ({<rVarName>} <type>) {`
    - first-class citizens
    - variadic `...`
    - closures

- Testing & Benchmarking
  - Write test cases in the same directory as source code
    - Test cases for `hello.go` will be in `hello_test.go`
  - `import "testing"`
  - `go test`
    - `<file>_test.go`
    - `func Test<>(t *testing.T)`
      - `t.Fail`
      - `t.Fatal`
      - `t.Skip`
  - `go test --bench . --benchmem`
    - `func Benchmark<>(b *testing.B)`
    - `b.N`
      - Don't use it as an argument for your code being benchmarked
      - Use it to run the code `N` times
  - Test Main
    - `func TestMain(m *testing.M)`
    - `m.Run()`

- Working with third-party packages
  - `go get`
    - `go get -u all`
  - `go mod`
    - `go mod init`
    - `go mod tidy`
    - `go mod why`
    - `go mod graph`
    - `go mod download`
  - `go list -m all`

- Packaging & Imports
  - `go help importpath`

- Methods
  - Receiver function `func (<varName> <receiverType>) <funcNameA>() {`
  - Pointer Receiver `func (<varName> *<receiverType>) <funcNameB>() {`
  - `var v <receiverType>`
    - `v.funcNameA()` & `v.funcNameB()`
  - `p := &v`
    - `p.funcNameA()` & `p.funcNameB()`

- Interfaces
  - `type <interfaceNamer> interface {`
  - implicit implementation
  - nil interface
  - empty interface

- Errors
  - `error` interface
    - `type error interface { Error() string }`
  - package [`errors`](https://golang.org/pkg/errors/)
    - `errors.New(<message>)`
    - `Unwrap() error`
    - `errors.Is` & `errors.As`

- Inheritance
  - `type Student struct { Person }`
  - `type ReadWriter interface { io.Reader ... }`
  - Multiple Inheritance
