# Web development - Refresher

- Writing our first "Hello, World!" API
  - `func <handlerName>(w http.ResponseWriter, req *http.Request) {`
  - `net/http`
    - `http.HandleFunc(<pathString>, <handlerFunc>)`
    - `http.ListenAndServe(":<PORT>", nil)`
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
