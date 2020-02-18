# Advanced Testing - Outline

- Coverage
  - `go test -coverprofile=coverage.out`
  - `go tool cover -html=coverage.out`

- Benchmarking [Comparison](https://godoc.org/golang.org/x/tools/cmd/benchcmp)
  - `go get golang.org/x/tools/cmd/benchcmp`
  - `go test -run=NONE -bench=. ./...`
  - `benchcmp old.txt new.txt`

- Stubbing in Go

- Mocking & Faking
  - [`gomock`](https://github.com/golang/mock)
  - [`faker`](https://github.com/bxcodec/faker)

- Writing bdd style unit tests
  - Setting up `ginkgo`
  - Using `gomega` matchers

- HTTP testing
  - `net/http/httptest` for dispatching requests to controller
  - `jarcoal/httpmock` for mocking external http requests

- Writing Cucumber in Go with [Godog](https://github.com/cucumber/godog)

---

## References

- [Gingko](https://onsi.github.io/ginkgo/) & [Gomega](https://onsi.github.io/gomega/)
- [Mocking HTTP Requests in Go](https://github.com/jarcoal/httpmock)
