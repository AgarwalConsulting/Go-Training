# Web development - Refresher

- Writing our first "Hello, World!" API
  - `net/http`
    - `http.ListenAndServe("<HOSTNAME>:<PORT>", nil)`
    - `mux := http.NewServeMux()`
    - `http.HandleFunc(<pathString>, <handlerFunc>)`
  - `func <handlerName>(w http.ResponseWriter, req *http.Request) {`
  - `encoding/json`
    - `json.NewEncoder(w).Encode(<varToBeSerialized>)`
    - `json.NewDecoder(req.Body).Decode(&v)`

- `gorilla/mux`
  - `r := mux.NewRouter()`
  - `r.HandleFunc(<path>, <handlerFunc>).Methods(<httpMethods>)`
  - `http.ListenAndServe(":<PORT>", r)`
  - Processing variables
    - `GET /people/{id}`
    - `mux.Vars(req)` returns `map[string]string`
      - Eg: `GET /people/1`
        - `mux.Vars(req)["id"] // "1"`

- `negroni` for middleware

- Basic Auth
  - `req.BasicAuth()`
