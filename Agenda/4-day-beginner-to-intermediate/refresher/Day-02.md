# Day 02

- Syntax
  - Variables, packages and functions
    - More Types (continued)
      - Arrays `var <name> [<length>]int`
      - Slices `var <name> []int`
        - `len` & `cap`
        - `append`
      - Ranges `for <index>, <value> := range <iterable> {`
      - Functions `func <funcName>(<argName> <type>) ({<rVarName>} <type>) {`
      - Maps `map[<keyType>]<valType>`
  - Methods and interfaces
    - Receiver function `func (<varName> <receiverType>) <funcName>() {`

- Testing & Benchmarking
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
  - `go mod`
    - `go mod init`
    - `go mod verify`
