# Day 03

- Writing our first 'Hello, World!' API
  - Introducing net/http package
  - Overview of Go standard library
    - fmt
    - io
    - os
    - net/http
    - encoding/json
  - `struct` tags
    - `json:"-"`

- Syntax overview!
  - Concurrency
    - Goroutines
      - `GOMAXPROCS`
      - Go scheduler
    - Channels
    - Select

- Introducing `mux`/`gin`, `negroni` & `gorilla` libraries for writing ReSTful APIs

- Structuring your application in Go using Clean Architecture
  - [Directory layout](https://github.com/golang-standards/project-layout)
  - [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
    - In Go
      - https://hackernoon.com/golang-clean-archithecture-efd6d7c43047
      - https://medium.com/@eminetto/clean-architecture-using-golang-b63587aa5e3f
        - [Code](https://github.com/eminetto/clean-architecture-go)

- 12 factor apps [https://12factor.net/] (Sample App: https://github.com/algogrit/yaes-server)

---

## References

- [Measuring context switching and memory overheads for Linux threads](https://eli.thegreenplace.net/2018/measuring-context-switching-and-memory-overheads-for-linux-threads/)
