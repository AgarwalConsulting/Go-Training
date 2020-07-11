# Basics - Refresher

- Goroutines: `go` keyword

- `go build` & `go run` commands
  - `GOOS` & `GOARCH` env variables

- `GOMAXPROCS`
  - Default value: number of CPU cores

- `sync.WaitGroup`

- Syntax
  - Packages
    - Every `.go` file starts with `package <name>`
    - All `.go` files in a directory need to belong to the same package
    - Every package needs to be in a directory of the same name
      - Except `main` package
    - `func main` needs to be defined in the `main` package

  - Exports
    - Works using the first character -> uppercase means exported

  - Variables, packages and functions
    - Variables
      - `var <name> <type>`
      - `var <name> = <val>`
      - `<name> := <val>`
    - zero values

    - Functions
      - Multiple returns
        `func hello() (int, string) {}`
      - named returns

    - Types
      - type conversion
        - `<type>(<val-of-other-type>)`
          - Eg: `val := int(42.00)`
        - Specialized type conversion `strconv`

    - constants:
      - `const` keyword
      - `iota` built-in

    - User defined types
      - `type` keyword
      - `type <NameOfUserDefinedType> <type>`

    - Flow control
      - for `for {`
        - normal `for <init>; <condition>; <post> {`
        - while `for <condition> {`
        - forever `for {`
          - can control execution using `break` & `continue`
        - `for` along with `range`
          - range returns `index` & `element`
          - `for index, element := range "abc" {`
      - if / else if / else `if <statement>; <condition> {`
      - switch `switch {`
        - normal `switch <value> { case <match>`
        - takes statement `switch <statement>; <value> { case <match>`
        - conditional `switch { case <bool-statement>: `
          - `if-else-if`
      - defers `defer <funcCall>`
        - stacked execution - FILO
        - Clean up with defer
        - Args will be evaluated before deferring
        - defer executes even in case of panic

    - More Types
      - Pointers `p := &<variable>`
      - Structs `type <name> struct {}`
      - Arrays `var <name> [<length>]int`
        - `len` & `cap`
      - Slices `var <name> []int`
        - References to underlying array
        - `append`
          - Eg. `append(s, <element1>, <element2>, ...)`
          - You have to assign it back to `s` or whatever is your slice variable name
          - Can have a different behavior depending on capacity
        - `make([]<type>, len, cap)`
        - `make([]<type>, len)`
      - Ranges `for <index>, <value> := range <iterable> {`
      - Maps `map[<keyType>]<valType>`
        - `make(map[string]Person)`
        - `len` only
        - `delete`
          - Eg. `delete(m, <key>)`
        - lookup returns 2 values - `val, ok`
