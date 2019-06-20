Day 02
------
- Syntax
  - More Types (continued)
    - Ranges `for <index>, <value> := range <iterable> {`
    - Functions `func <funcName>(<argName> <type>) ({<rVarName>} <type>) {`
    - Maps `map[<keyType>]<valType>`
  - Methods and interfaces
    - Receiver function `func (<varName> <receiverType>) <funcName>() {`
    - interfaces `type <interfaceNamer> interface {`

- Writing our first "Hello, World!" API
  - `func <handlerName>(w http.ResponseWriter, req *http.Request) {`
  - `net/http`
    - `http.HandleFunc(<pathString>, <handlerFunc>)`
    - `http.ListenAndServe(":<PORT>", nil)`
  - `encoding/json`
    - `json.NewEncoder(w).Encode(<varToBeSerialized>)`

- Working with third-party packages
  - `dep`
    - Setup: http://github.com/Chennai-Golang/101-workshop/blob/master/Setup.md
    - `Gopkg.toml`, `Gopkg.lock` & `vendor`
  - `gorilla/mux`
    - `r := mux.NewRouter()`
    - `r.HandleFunc(<path>, <handlerFunc>).Methods(<httpMethods>)`
    - `http.ListenAndServe(":<PORT>", r)`
    - Processing variables
  - `negroni`

- Working with relational databases
  - `database/sql`
  - `jinzhu/gorm`

- Misc
  - Gin (https://github.com/gin-gonic/gin)
  - Error handling
    - https://dave.cheney.net/tag/error-handling
    - https://blog.golang.org/errors-are-values
  - Hot Reloading: http://github.com/codegangsta/gin
  - `struct` tags
    - `json:"-"`
