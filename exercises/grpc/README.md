# Biblioteca gRPC

Revisiting the Biblioteca ReST API, but now as a gRPC based API.

## Setup

### Install protobuf for your platform

#### For Mac

```bash
brew install protobuf
```

#### For Linux

* From root of `Go-Training`, run:

```bash
./scripts/setup-linux.sh
```

#### For Windows

* From protobuf [release page](https://github.com/protocolbuffers/protobuf/releases`)
  * Download `protoc-<latest_version>-win64.zip`
    * E.g. latest_version=3.12.3, download `protoc-3.12.3-win64.zip`
* Create a folder `C:/protobuf`
  * Copy the zip into `C:/protobuf`
  * Extract the zip into `C:/protobuf` folder
* Add the path `C:/protobuf/protoc-<latest_version>-win64/bin` in to your `%PATH%` variable

### Install go gRPC plugin

```bash
go get github.com/golang/protobuf/protoc-gen-go
```

## Development

* Update the `biblioteca.proto` file
* Update the generated stubs using:

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/biblioteca.proto
```

## Instructions

* Define new / Modify `rpc` & `message`s in `biblioteca.proto` file
  * Implement at least one streaming rpc apart from simple unary rpcs
* Regenerate the `biblioteca.pb.go` using the command in `scripts/generate.sh`
* Implement the server
* Generate clients stubs for a language of your choice, you can use one of: Java, Python, Ruby, JavaScript or any other language
* Create a client, written in language of your choice, communicating with the server, written in Go
* Write a logging middleware, or as they are called in gRPC `interceptors`, using:
  * `grpc.UnaryInterceptor(...)`
  * `grpc.StreamInterceptor(...)`
* Use [`go-grpc-middleware`](https://github.com/grpc-ecosystem/go-grpc-middleware) to write chain interceptors
