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

- Writing our first "Hello, World!" API
  - `func <handlerName>(w http.ResponseWriter, req *http.Request) {`
  - `net/http`
    - `http.HandleFunc(<pathString>, <handlerFunc>)`
    - `http.ListenAndServe(":<PORT>", nil)`
  - `encoding/json`
    - `json.NewEncoder(w).Encode(<varToBeSerialized>)`

- Working with third-party packages
  - `gorilla/mux`
    - `r := mux.NewRouter()`
    - `r.HandleFunc(<path>, <handlerFunc>).Methods(<httpMethods>)`
    - `http.ListenAndServe(":<PORT>", r)`
    - Processing variables
      - `GET /people/{id}`
      - `mux.Vars()`
