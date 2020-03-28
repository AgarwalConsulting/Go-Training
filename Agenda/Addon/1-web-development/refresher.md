# Web development - Refresher

- Writing our first "Hello, World!" API
  - `net/http`
    - `http.ListenAndServe(":<PORT>", nil)`
    - `mux := http.NewServeMux()`
    - `http.HandleFunc(<pathString>, <handlerFunc>)`
  - `func <handlerName>(w http.ResponseWriter, req *http.Request) {`
  - `encoding/json`
    - `json.NewEncoder(w).Encode(<varToBeSerialized>)`

- `gorilla/mux`
  - `r := mux.NewRouter()`
  - `r.HandleFunc(<path>, <handlerFunc>).Methods(<httpMethods>)`
  - `http.ListenAndServe(":<PORT>", r)`
  - Processing variables
    - `GET /people/{id}`
    - `mux.Vars()`

- `negroni` for middleware
