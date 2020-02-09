# Performance tuning, instrumentation & debugging - Outline

- Profiling
  - Run one profiler at a time, CPU profiler takes measurements every 10ms
  - `go test -bench=.`
  - `go test -cpuprofile cpu.prof -memprofile mem.prof -bench .`
  - `go tool pprof cpu.prof`
  - `go tool pprof --alloc_objects mep.prof`
  - `go tool pprof --inuse_objects mep.prof`
  - Profiling web apps
    - `import _ "net/http/pprof"`
  - `go tool pprof -seconds 10 http://localhost:8080/debug/pprof/profile`
    - `wrk -t 4 -c 16 -d 30 http://localhost:8000/posts`
  - `go get -u github.com/google/pprof`
    - `pprof cpu.prof`
    - `pprof -seconds 10 -http=localhost:8181 http://localhost:8080/debug/pprof/profile`

- Go runtime
  - Garbage collector
  - Scheduler

- [Tracing](https://golang.org/pkg/runtime/trace/)
  - [HTTP Tracing](https://golang.org/pkg/net/http/httptrace/)

- Debugging using dlv

- Distributed Tracing
  - [Jaeger](https://github.com/jaegertracing/jaeger-client-go)

---

## References

- [Profiling Go Programs - Go blog](https://blog.golang.org/profiling-go-programs)
- [Profiling in Go by Cory Lanou - Go Study Group](https://www.youtube.com/watch?v=YNye3SZWvj8)
- [Gopher Guides](https://www.gopherguides.com/courses/advanced/modules/advanced-profiling/)
- [Profiling Go programs with pprof by Julia Evans](https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/)
- [An Introduction to go tool trace](https://about.sourcegraph.com/go/an-introduction-to-go-tool-trace-rhys-hiltner)
- [Optimizing Go Programs](https://github.com/sathishvj/optimizing-go-programs)
- [Introducing HTTP Tracing](https://blog.golang.org/http-tracing)
- [Debugging what you deploy in Go](https://blog.golang.org/debugging-what-you-deploy)
