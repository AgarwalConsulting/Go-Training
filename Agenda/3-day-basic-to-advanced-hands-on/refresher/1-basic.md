# Basics - Refresher

- Goroutines: `go` keyword

- `go build` & `go run` commands

- `GOMAXPROCS`

- `sync.WaitGroup`

- Syntax
  - Variables, packages and functions
    - Variables
      - `var <name> <type>`
      - `var <name> = <val>`
      - `<name> := <val>`

    - Functions
      - Multiple returns
        `func hello() (int, string) {}`
      - named returns, zero values

    - Types
      - type coercion
        - `<type>(<val-of-other-type>)`
          - Eg: `val := int(42.00)`
      - type conversion
        - `strconv`
      - inference
        - any type
          - `interface{}`

    - constants:
      - `const` keyword
      - `iota` built-in

    - Flow control
      - for `for {`
        - normal `for <init>; <condition>; <post> {`
        - while `for <condition> {`
        - forever `for {`
          - can control execution using `break`
        - `for` along with `range`
      - if / else if / else `if <statement>; <condition> {`
      - switch `switch {`
        - normal `switch <value> {`
        - conditional `switch <value> { case <statement>: `
        - no condition switch `if-else-if`
      - defers `defer <funcCall>`

    - More Types
      - Pointers `p := &<variable>`
      - Structs `type <name> struct {}`
      - Arrays `var <name> [<length>]int`
        - `len` & `cap`
      - Slices `var <name> []int`
        - References to underlying array
        - `append`
          - Can have a different behavior depending on capacity
        - `make([]<type>, len, cap)`
        - `make([]<type>, len)`
      - Ranges `for <index>, <value> := range <iterable> {`
      - Maps `map[<keyType>]<valType>`
        - `make(map[string]Person)`
        - `len` only
        - `delete`
        - lookup returns 2 values - `val, ok`
