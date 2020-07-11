# Web development - Refresher

## ReSTful / http API

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

- `negroni` or `alice` for middleware

## gRPC

- Writing out first "Hello, World!" rpc
  - Start by defining your service & messages in a proto file, with `syntax = "proto3"`
  - Generate the client & server stubs, using `protoc`
- Implementing client
  - Connect to the server using address, `conn, err := grpc.Dial(address, ...opts)`
  - Create a new client from generated code, `pb.NewHelloServiceClient()`
- Implementing server
  - Create a tcp socket to listen on `lis, err := net.Listen("tcp", address)`
  - Create a new grpc server `s := grpc.NewServer()`
  - Register an implmentation of `pb.HelloServiceServer` using: `pb.RegisterHelloServiceServer(s, &HelloService{})`
  - Start listening on the tcp socket `s.Serve(lis)`

## Auth

- Basic Auth
  - `req.BasicAuth()`
