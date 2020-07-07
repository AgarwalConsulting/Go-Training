# Overview

## Sweets spots of Go

* Why Go?

### Compiling cross-platform binaries

* `GOOS` & `GOARCH` environment variables

### Concurrency vs Parallelism

OS Overview

![OS Overview](assets/1-OS-process-and-its-threads.png)

> Image Credits: https://github.com/sathishvj/optimizing-go-programs#m-p-g

Concurrency vs Parallelism

![Concurrency vs Parallelism](assets/4-concurrency-and-parallelism.png)

> Image Credits: https://github.com/sathishvj/optimizing-go-programs#m-p-g

* Overview of Go scheduler

OS vs Go Scheduler

![OS vs Go Scheduler](assets/go-vs-os-scheduler.png)

> Image Credits: https://speakerdeck.com/kavya719/the-scheduler-saga?slide=15

OS Thread & Go Routine Scheduling

![Goroutine vs OS Thread](assets/os-thread-goroutine.png)
> Image Credits: https://speakerdeck.com/kavya719/the-scheduler-saga?slide=18

* `GOMAXPROCS` controls the max number of OS threads a go program can create

### Other Packages

* `context` package
  * Useful for signaling multiple goroutines at the same time
  * Useful for controlling goroutine execution
  * Useful for sharing values across goroutine boundaries

---

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

## Deploying Go

### Understanding [12-factor apps](http://12factor.net/)

### Example app: [YAES Server](https://github.com/algogrit/yaes-server)

#### Multi-stage Docker builds

* Binary is all you need

#### Running & managing Go on k8s cluster

* [Nginx Ingress](https://github.com/kubernetes/ingress-nginx)
* [Cert Manager](https://cert-manager.io/docs)

---

## Observability

### What is Observability?

* Logging provides insight into application-specific messages emitted by processes.
* Metrics provide quantitative information about processes running inside the system, including counters, gauges, and histograms.
* Tracing, aka distributed tracing, provides insight into the full lifecycles, aka traces, of requests to the system, allowing you to pinpoint failures and performance issues.

### Metrics & Statistics

* Push vs Pull mechanisms

![Prometheus Scrape](assets/monitoring/prometheus-scrape.png)
> Image Credits: https://blog.pvincent.io/2017/12/prometheus-blog-series-part-3-exposing-and-collecting-metrics/

* Introspection using `/healthz`
  * [Liveness](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/) probes in K8s

### Prometheus & Grafana

![Prometheus Architecture](assets/monitoring/prometheus-architecture.png)
> Image Credits: https://logz.io/blog/prometheus-vs-graphite/

* Prometheus [handler for Go](https://prometheus.io/docs/guides/go-application/)
* Using sidecar pattern for [kubernetes](https://www.weave.works/blog/prometheus-and-kubernetes-monitoring-your-applications/)

![Prometheus Gathering](assets/monitoring/prometheus-gathering.png)
> Image Credits: https://devconnected.com/the-definitive-guide-to-prometheus-in-2019/

### [OpenTelemetry](https://opentelemetry.io/)

* What is distributed tracing?
* OpenCensus + OpenTracing = OpenTelemetry
* [Specification driven](https://github.com/open-telemetry/opentelemetry-specification)
* [Go Client](https://github.com/open-telemetry/opentelemetry-go/blob/master/README.md)

![Pluggable Architecture](assets/monitoring/opentelemetry-pluggable.png)
> Image Credits: https://medium.com/opentelemetry/opentelemetry-beyond-getting-started-5ac43cd0fe26
