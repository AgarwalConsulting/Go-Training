# Overview

## Testing & Benchmarking

### Testing

- Write test cases in the same directory as source code
  - Test cases for `hello.go` will be in `hello_test.go`
- `import "testing"`
- `go test`
  - `<file>_test.go`
  - `func Test<>(t *testing.T)`
    - `t.Fail`
    - `t.Fatal`
    - `t.Skip`
- Test Main
  - `func TestMain(m *testing.M)`
  - `m.Run()`

### Coverage

- `go test -coverprofile=coverage.out`
- `go tool cover -html=coverage.out`

### Benchmarking

- `go test --bench . --benchmem`
  - `func Benchmark<>(b *testing.B)`
  - `b.N`
    - Don't use it as an argument for your code being benchmarked
    - Use it to run the code `N` times
- Benchmarking [Comparison](https://godoc.org/golang.org/x/tools/cmd/benchcmp)
  - `go get golang.org/x/tools/cmd/benchcmp`
  - `go test -run=NONE -bench=. ./...`
  - `benchcmp old.txt new.txt`

## Structuring your code

- [Directory layout](https://github.com/golang-standards/project-layout)
- Writing modular code by adopting [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

## Advanced

### Stubbing in Go

### Mocking & Faking

- [`gomock`](https://github.com/golang/mock)
- [`faker`](https://github.com/bxcodec/faker)

### Writing bdd style unit tests

- `Rspec` or `JBehave` style
  - Setting up `ginkgo`
  - Using `gomega` matchers
- `Cucumber` style
  - Using [`godog`](https://github.com/cucumber/godog)

### HTTP testing

- `net/http/httptest` for dispatching requests to controller
- `jarcoal/httpmock` for mocking external http requests
