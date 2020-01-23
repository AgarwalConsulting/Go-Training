# Day 04

- Tour of Golang continued
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

- Packaging a Go app
  - Peeking into sample `Dockerfile`
  - Intoducing `make` command; peeking into sample `Makefile`
- Deploying a sample app to Heroku!

- Testing, Benchmarking and Profiling
  - Testing
    - xUnit style tests using built-in `testing` package (for Java)
    - BDD style tests using `ginkgo` & `gomega` (for Ruby)
    - Overview of `net/http/httptest`

- Benchmarking and profiling

# Assignment

- Building your own pipeline
