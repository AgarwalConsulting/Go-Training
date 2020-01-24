# Day 02

- Syntax
  - Variables, packages and functions
    - More Types
      - Pointers `p := &<variable>`
      - Structs `type <name> struct {}`
      - Arrays `var <name> [<length>]int`
      - Slices `var <name> []int`
        - `len` & `cap`
        - `append`
  - More Types (continued)
    - Ranges `for <index>, <value> := range <iterable> {`
    - Functions `func <funcName>(<argName> <type>) ({<rVarName>} <type>) {`
    - Maps `map[<keyType>]<valType>`
  - Methods and interfaces
    - Receiver function `func (<varName> <receiverType>) <funcName>() {`
    - interfaces `type <interfaceNamer> interface {`
    - errors
      - `error` is an interface
        - `type error interface { Error() string }`
      - package `errors`
        - `errors.New(<message>)`

- Inheritance
  - `type Student struct { Person }`
  - `type ReadWriter interface { io.Reader ... }`

- Writing our first "Hello, World!" API
  - `func <handlerName>(w http.ResponseWriter, req *http.Request) {`
  - `net/http`
    - `http.HandleFunc(<pathString>, <handlerFunc>)`
    - `http.ListenAndServe(":<PORT>", nil)`
  - `encoding/json`
    - `json.NewEncoder(w).Encode(<varToBeSerialized>)`
