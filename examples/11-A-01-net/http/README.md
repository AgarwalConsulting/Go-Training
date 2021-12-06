# Writing HTTP Applications `net/http`

## Starting a server

- `http.ListenAndServe(<addrString>, <router>)`

## Routing Requests

- `r := http.NewServeMux()`
  - `r.HandleFunc("url/pattern", handlerFunc)`

- handlerFunc => `func(w http.ResponseWriter, req *http.Request)`

## Creating Response

- `http.ResponseWriter` interface => Superset of `io.Writer`
  - `Write([]byte) (int, err)`
  - `Header()`
    - `Add(<res-header>, <res-header-val>)`
    - `Set(<res-header>, <res-header-val>)`

Writing a [http status code](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status)

- Different status codes:
  - 2xx (ok)
  - 3xx (redirects => http://google.com => https://google.com)
  - 4xx (client error)
  - 5xx (server error)

- `WriteHeader(int)`
  - Default Status code: 200

### Responding with JSON

- `encoding/json`
  - `NewEncoder` => serializing go DS to JSON
  - `NewDecoder` => deserializing JSON to go DS

  - `Marshler` & `Unmarshaler` interfaces
  - Alternative `json` struct tags

## Dealing with request

- `*http.Request` => struct
  - `Body`
  - `Method`

## Third-party router

- `gorilla/mux`
  - `r := mux.NewRouter()`

Allows chaining:

- `r.HandleFunc("url/pattern", handlerFunc).Methods("HTTPMEthod")`
