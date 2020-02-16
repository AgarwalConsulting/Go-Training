# Performance tuning, instrumentation, debugging & analysis - Outline

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

- Escape analysis
  - `go build -gcflags "-m -m"`
  - Understanding Stack and Heap memory

---

## References

- [Profiling Go Programs - Go blog](https://blog.golang.org/profiling-go-programs)
- [(Video) Profiling in Go by Cory LaNou - Go Study Group](https://www.youtube.com/watch?v=YNye3SZWvj8)
- [Gopher Guides](https://www.gopherguides.com/courses/advanced/modules/advanced-profiling/)
- [Profiling Go programs with pprof by Julia Evans](https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/)
- [Continuous Profiling of Go programs](https://medium.com/google-cloud/continuous-profiling-of-go-programs-96d4416af77b)
- [An Introduction to go tool trace](https://about.sourcegraph.com/go/an-introduction-to-go-tool-trace-rhys-hiltner)
- [Optimizing Go Programs](https://github.com/sathishvj/optimizing-go-programs)
- [Introducing HTTP Tracing](https://blog.golang.org/http-tracing)
- [Debugging what you deploy in Go](https://blog.golang.org/debugging-what-you-deploy)
- [Language Mechanics On Escape Analysis](https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-escape-analysis.html)
- [Escape Analysis Flaws](https://www.ardanlabs.com/blog/2018/01/escape-analysis-flaws.html)
- [Design Philosophy On Data And Semantics](https://www.ardanlabs.com/blog/2017/06/design-philosophy-on-data-and-semantics.html)
- [(Video) GopherCon 2018: Kavya Joshi - The Scheduler Saga](https://www.youtube.com/watch?v=YHRO5WQGh0k)
- [(Video) Maya Rosecrance - The garbage collector](https://www.youtube.com/watch?v=qj5a4ZEsttg)
- [(Video) Golang's Realtime Garbage Collector](https://pusher.com/sessions/meetup/the-realtime-guild/golangs-realtime-garbage-collector)
- [Go's work-stealing scheduler](https://rakyll.org/scheduler/)
