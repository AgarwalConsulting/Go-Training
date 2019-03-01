Day 03
------

- `go help importpath`
- Go concurrency patterns (https://www.youtube.com/watch?v=f6kdp27TYZs)
  - Producer-consumer
  - Futures (similar to `async`-`await`)
    - Map Reduce
  - Generators
  - Pipeline example

- Testing, Benchmarking and Profiling
  - Testing
    - xUnit style tests using built-in `testing` package (for Java)
    - BDD style tests using `ginkgo` & `gomega` (for Ruby)
    - Overview of `net/http/httptest`

- Benchmarking and profiling (https://www.youtube.com/watch?v=YNye3SZWvj8)
  - https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
  - https://blog.golang.org/profiling-go-programs
  - https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/
  - https://www.youtube.com/watch?v=YNye3SZWvj8
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
