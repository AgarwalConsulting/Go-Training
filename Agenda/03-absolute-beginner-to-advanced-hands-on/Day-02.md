# Day 02

- Syntax overview!
  - Going through 'Tour of Go' site:
    - More types
      - Pointers
      - Struct Fields
      - Imports / Exports
      - Pointers & structs
      - Arrays
      - Slices (https://gobyexample.com/slices)
        - how they affect the underlying array
        - `len` and `capacity`
      - Ranges (https://gobyexample.com/range)
      - Functions (https://golang.org/doc/codewalk/functions/)
    - Methods and interfaces
      - Methods and pointer indirection
      - Receiver functions
        - pointer receivers vs value receivers
      - interfaces
        - implicit implementation
        - empty interface
    - `init` function
    - Recovering from a panic using `recover` (https://blog.golang.org/defer-panic-and-recover)
  - Concurrency
    - Goroutines
    - Channels
    - Select
- Setting up `dep` and learning about the `bin` directory
  - Usage of `dep` as a package manager
- Introducing `mux`/`gin`, `negroni` & `gorilla` libraries for writing ReSTful APIs
- Writing our first 'Hello, World!' API
  - Introducing net/http package
  - Overview of Go standard library
    - fmt
    - io
    - os
    - net/http
    - encoding/json

## Assignment

- Writing a simple "Hello, World" API using native packages
- Writing a simple "Hello, World" API using gin/mux
