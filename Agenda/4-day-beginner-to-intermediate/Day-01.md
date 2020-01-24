# Day 01

- Why Go?
- Introduce Goroutines (Power of Go)
  - GOMAXPROCS
  - sync.WaitGroup
  - Go scheduler
  - Deadlock detection
- `go build` and `go run`
  - binaries (`GOOS` & `GOARCH` / `CGO_ENABLED`)
- Setting up Go environment: $GOPATH & $PATH
  - Including setting up VSCode
  - Setting up delve debugger (optional)
  - Check using `101-workshop` repo [README.md](https://github.com/Chennai-Golang/101-workshop)
- Syntax overview!
  - Going through 'Tour of Go' site:
    - Variables, packages and functions
      - Multiple Returns
      - Named return values
        - `defer`
      - Zero values
      - Type conversions
      - Type inference
      - Constants
        - Constants in standard library
        - Used as [enums](https://golang.org/pkg/time/#pkg-constants)
    - Flow control
      - For
        - as `while`
        - range
        - forever
      - if
        - with [statements](https://tour.golang.org/flowcontrol/6)
      - defers
        - [stack based execution](https://tour.golang.org/flowcontrol/13)
      - [Switch statements](https://gobyexample.com/switch)
        - no break
        - statements are optional
        - top-down execution - where execution stops after a case succeeds

## Assignment

- Write a simple fizz-buzz type program
