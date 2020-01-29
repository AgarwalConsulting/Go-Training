# Day 02

- Syntax
  - Variables, packages and functions
    - More Types (continued)
      - Arrays `var <name> [<length>]int`
        - `len` & `cap`
      - Slices `var <name> []int`
        - `append`
        - `make([]<type>, len, cap)`
        - `make([]<type>, len)`
      - Ranges `for <index>, <value> := range <iterable> {`
      - Functions `func <funcName>(<argName> <type>) ({<rVarName>} <type>) {`
      - Maps `map[<keyType>]<valType>`
        - `make(map[string]Person)`
  - Methods and interfaces
    - Receiver function `func (<varName> <receiverType>) <funcNameA>() {`
    - Pointer Receiver `func (<varName> *<receiverType>) <funcNameB>() {`
    - `var v <receiverType>`
      - `v.funcNameA()` & `v.funcNameB()`
    - `p := &v`
      - `p.funcNameA()` & `p.funcNameB()`

- Testing & Benchmarking
  - `import "testing"`
  - `go test`
    - `<file>_test.go`
    - `func Test<>(t *testing.T)`
  - `go test --bench . --benchmem`
    - `func Bench<>(b *testing.B)`
    - `b.N`
  - Test Main
    - `func TestMain(m *testing.M)`
    - `m.Run()`

- Working with third-party packages
  - `go get`
  - `go mod`
    - `go mod init`
