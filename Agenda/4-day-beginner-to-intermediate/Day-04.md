# Day 04

- Testing & Benchmarking
  - xUnit style tests using built-in `testing` package (for Java)
  - BDD style tests using `ginkgo` & `gomega` (for Ruby)
  - Overview of `net/http/httptest`
  - Mocking in Go

- Other built-in packages
  - [`reflect`](https://golang.org/pkg/reflect/)
  - [`context`](https://golang.org/pkg/context/)
  - [`flag`](https://golang.org/pkg/flag/)
  - [`os`](https://golang.org/pkg/os/)

- [Awesome](https://github.com/avelino/awesome-go) 3rd party packages
  - [`validator`](https://gopkg.in/go-playground/validator.v10)
  - [`viper`](https://github.com/spf13/viper)
  - [`gin`](https://github.com/codegangsta/gin) - for hot reloading

- Concurrency patterns

- Looking at a sample app [`go-auction-api`](https://github.com/algogrit/go-auction-api)
  - config
  - logging
  - Error reporting
  - JWT token

- Packaging a Go app
  - Peeking into sample `Dockerfile`

- Profiling
  - Run one profiler at a time, CPU profiler takes measurements every 10ms
  - `go test -bench=.`
  - `go test -cpuprofile cpu.prof -memprofile mem.prof -bench .`
  - `go tool pprof cpu.prof`
  - `go tool pprof --alloc_objects mep.prof`
  - `go tool pprof --inuse_objects mep.prof`
  - `import _ "net/http/pprof"`
  - `go tool pprof -seconds 10 http://localhost:8080/debug/pprof/profile`
    - `wrk -t 4 -c 16 -d 30 http://localhost:8000/posts`
  - `go get -u github.com/google/pprof`
    - `pprof cpu.prof`
    - `pprof -seconds 10 -http=localhost:8181 http://localhost:8080/debug/pprof/profile`

- Tracing

- Go runtime
  - Garbage collector
  - Scheduler

---

## References

- [How to write benchmarks in Go by Dave Cheney](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)
- [Profiling Go Programs - Go blog](https://blog.golang.org/profiling-go-programs)
- [Profiling in Go by Cory Lanou - Go Study Group](https://www.youtube.com/watch?v=YNye3SZWvj8)
- [Profiling Go programs with pprof by Julia Evans](https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/)
