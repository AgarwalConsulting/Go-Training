# Basics

- Why Go?

- [Setting up Go](https://github.com/Chennai-Golang/101-workshop/blob/master/Setup.md) environment: `$GOPATH` & `$PATH`
  - Including setting up VSCode
  - Setting up delve debugger (optional)
  - Check using `101-workshop` repo [README.md](https://github.com/Chennai-Golang/101-workshop)

- Introduce Goroutines (Power of Go)
  - sync.WaitGroup
  - GOMAXPROCS
  - Go scheduler

- `go build` and `go run`
  - binaries (`GOOS` & `GOARCH` / `CGO_ENABLED`)

- Syntax overview!
  - Going through 'Tour of Go' site
    - Variables, packages and functions
      - Multiple Returns
      - Named return values
      - Zero values
      - Type conversions
      - Type inference
      - Constants
        - Constants in standard library
        - Used like [enums](https://golang.org/pkg/time/#pkg-constants)

    - Flow control
      - For
        - as `while`
        - range
        - forever
      - if
        - with [statements](https://tour.golang.org/flowcontrol/6)
      - [Switch statements](https://gobyexample.com/switch)
        - no break
        - statements are optional
        - top-down execution - where execution stops after a case succeeds
      - [Ranges](https://gobyexample.com/range)
      - `defer`
        - [stack based execution](https://tour.golang.org/flowcontrol/13)

    - More types
      - Pointers
      - Struct Fields
      - Imports / Exports
      - Pointers & structs
      - Arrays
      - [Slices](https://blog.golang.org/slices-intro)
        - how they affect the underlying array
        - `len` and `capacity`
      - Maps
        - `map`
        - `len` only
        - lookup, insert & delete

---

## Assignment

- First 4 basic assignments

## References

- [Go's Declaration Syntax](https://blog.golang.org/gos-declaration-syntax)
- [GOOS & GOARCH values](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63)
