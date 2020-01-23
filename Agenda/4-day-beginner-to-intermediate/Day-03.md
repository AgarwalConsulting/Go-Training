# Day 03

- Syntax overview!
  - Going through 'Tour of Go' site:
    - Methods and interfaces
      - Methods and pointer indirection
      - Receiver functions
        - pointer receivers vs value receivers
      - interfaces
        - implicit implementation
        - empty interface
    - `init` function
  - Concurrency
    - Goroutines
    - Channels
    - Select

- `init` function
- 12 factor apps [https://12factor.net/] (Sample App: https://github.com/algogrit/yaes-server)
- `struct` tags
  - `json:"-"`

- Structuring your application in Go using Clean Architecture
  - Project Layout
    - https://medium.com/@eminetto/clean-architecture-using-golang-b63587aa5e3f
      - https://github.com/eminetto/clean-architecture-go
    - https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
    - https://github.com/golang-standards/project-layout
    - https://www.reddit.com/r/golang/comments/ay1sdf/best_way_to_learn_proper_go_design_patterns_and/

- Looking at a sample app (`go-auction-api`)
  - config
  - logging
  - Error reporting
  - JWT token
