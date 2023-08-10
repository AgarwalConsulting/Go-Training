# Overview

## Performance Engineering

### Writing benchmarks

* `go test --bench . --benchmem`
  * `func Benchmark<>(b *testing.B)`
  * `b.N`
    * Don't use it as an argument for your code being benchmarked
    * Use it to run the code `N` times

### API Load Testing

* Common across languages
  * `wrk` for ReSTful/SOAP APIs
    * `wrk -t 4 -c 16 -d 30 <url>`
  * `ghz` for gRPC

### Profiling

* CPU & Memory
  * Run one profiler at a time, CPU profiler takes measurements every 10ms
  * `go test -cpuprofile cpu.prof -memprofile mem.prof -bench .`
    * `go tool pprof cpu.prof`
    * `go tool pprof --alloc_objects mep.prof`
    * `go tool pprof --inuse_objects mep.prof`
  * Installation
    * `go get -u github.com/google/pprof`
  * Profiling [web apps](https://golang.org/pkg/net/http/pprof/)
    * `import _ "net/http/pprof"`
    * `pprof -seconds 10 -http=localhost:8181 http://localhost:8080/debug/pprof/profile`
* Block
  * *shows where goroutines block waiting on synchronization primitives*
  * `runtime.SetBlockProfileRate`
* Mutex
  * *reports the lock contentions*
  * `runtime.SetMutexProfileFraction`
* Tracing
  * `go test -trace <trace_file>`
    * `go tool trace <trace_file>`
  * `runtime/trace` package
  * `golang.org/x/net/trace` package
  * [HTTP tracing](https://blog.golang.org/http-tracing)
    * `net/http/httptrace` package

* [Runtime statistics & events](https://golang.org/doc/diagnostics.html#godebug)
*Runtime also emits events and information if GODEBUG environmental variable is set accordingly.*

  * `GODEBUG=gctrace=1`
  * `GODEBUG=schedtrace=X`

---

## Building and deploying more optimised code

- [Using PGO](https://go.dev/doc/pgo)

---

## [Observability](https://github.com/AgarwalConsulting/Kubernetes-Training/blob/master/notes/observability.md)
