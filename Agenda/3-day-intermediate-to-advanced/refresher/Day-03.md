Day 03
------

- Concurrency
  - Goroutines
    - `go` keyword
    - M:N scheduler
  - sync
    - sync.WaitGroup
    - sync.Mutex
  - Channels
    - Buffered `make(chan <type>, <n>)`
    - Unbuffered `make(chan <type>)`
      - `<-ch`
        - Receiving from a channel; blocks if there is nothing to receive
      - `ch <-`
        - Sending to a channel;
          - blocks if the buffer is full! (Buffered channel)
          - blocks until there is a receive! (Unbuffered channel)
  - Select
    - The select statement lets a goroutine wait on multiple communication operations.
    - A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
    - The default case in a select is run if no other case is ready.

- Misc
  - `init` function
  - 12 factor apps [https://12factor.net/] (Sample App: https://github.com/algogrit/yaes-server)
  - Go project layout (https://github.com/golang-standards/project-layout)

- Project Layout (https://github.com/AgarwalConsulting/go-auction-api)
  - https://medium.com/@eminetto/clean-architecture-using-golang-b63587aa5e3f
    - https://github.com/eminetto/clean-architecture-go
  - https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
  - https://github.com/golang-standards/project-layout

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

- Packaging a Go app
  - Multi-stage docker builds
