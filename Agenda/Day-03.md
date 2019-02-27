Day 03
------

- Testing, Benchmarking and Profiling
  - Testing
    - xUnit style tests using built-in `testing` package (for Java)
    - BDD style tests using `ginkgo` & `gomega` (for Ruby)
    - Overview of `net/http/httptest`
  - Benchmarking (https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)
  - Profiling (https://blog.golang.org/profiling-go-programs)
    - `go test -bench=.`
    - `go test -cpuprofile cpu.prof -memprofile mem.prof -bench .`
    - `import _ "net/http/pprof"`
- Packaging a Go app
  - Peeking into sample `Dockerfile`
  - Intoducing `make` command; peeking into sample `Makefile`
- Deploying a sample app to Heroku!
