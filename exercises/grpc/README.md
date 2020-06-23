# Biblioteca gRPC

Revisiting the Biblioteca ReST API, but now as a gRPC based API.

## Instructions

* Define new / Modify `rpc` & `message`s in `biblioteca.proto` file
  * Implement at least one streaming rpc apart from simple unary rpcs
* Regenerate the `biblioteca.pb.go` using the command in `scripts/generate.sh`
* Implement the server
* Generate clients stubs for a language of your choice, you can use one of: Java, Python, Ruby, JavaScript or any other language
* Create a client, written in language of your choice, communicating with the server, written in Go
