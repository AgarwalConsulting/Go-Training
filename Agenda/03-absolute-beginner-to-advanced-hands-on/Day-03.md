# Day 03

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
- Working with third-party packages
  - `dep`
    - Setup: http://github.com/Chennai-Golang/101-workshop/blob/master/Setup.md
    - `Gopkg.toml`, `Gopkg.lock` & `vendor`
  - `gorilla/mux`
    - `r := mux.NewRouter()`
    - `r.HandleFunc(<path>, <handlerFunc>).Methods(<httpMethods>)`
    - `http.ListenAndServe(":<PORT>", r)`
    - Processing variables
  - `negroni`

- Working with relational databases
  - `database/sql`
  - `jinzhu/gorm`
  - `jmoiron/sqlx`

- `init` function
- 12 factor apps [https://12factor.net/] (Sample App: https://github.com/algogrit/yaes-server)
- `struct` tags
  - `json:"-"`

- `reflection` package

## Assignment

- Exercising on exercism.io
