# Assignment

To start off, install mux:

```bash
go get -u github.com/golang/dep/cmd/dep
dep ensure
```

## Run

```bash
go build .

./assignment
```

This starts a server on port `9000`.

## Test

```bash
./test-with-curl.sh
```
