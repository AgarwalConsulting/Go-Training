# Day 03

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

- Packaging a Go app
  - Multi-stage docker builds
