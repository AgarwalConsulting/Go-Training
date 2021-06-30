#!/usr/bin/env bash

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/biblioteca.proto
# protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative api/biblioteca.proto

# protoc --plugin=protoc-gen-grpc-java --proto_path=api "biblioteca.proto" --grpc-java_out="src"
