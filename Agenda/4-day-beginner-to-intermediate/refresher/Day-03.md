# Day 03

- Syntax

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

- Working with third-party packages
  - `dep`
    - Setup: http://github.com/Chennai-Golang/101-workshop/blob/master/Setup.md
    - `Gopkg.toml`, `Gopkg.lock` & `vendor`
  - `gorilla/mux`
    - `r := mux.NewRouter()`
    - `r.HandleFunc(<path>, <handlerFunc>).Methods(<httpMethods>)`
    - `http.ListenAndServe(":<PORT>", r)`
    - Processing variables
      - `GET /people/{id}`
      - `mux.Vars()`
  - `negroni`

- Misc
  - Error handling
    - https://dave.cheney.net/tag/error-handling
    - https://blog.golang.org/errors-are-values
  - `struct` tags
    - `json:"-"`
