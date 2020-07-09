# Fibonacci gRPC Example

This server implements both `unary` and `streaming` rpc calls. At the core is the `generator/fibonacci.go`, which is a "goroutine-safe" fibonacci generator.

You can find the generated stubs & proto definition in `api/fibonacci.proto` directory.

## Server

You can start the server using:

```bash
go run ./cmd/server
```

## Client

You can start the client using:

```bash
go run ./cmd/client
```
