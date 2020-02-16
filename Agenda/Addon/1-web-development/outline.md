# Web development - Outline

- Writing our first 'Hello, World!' API
  - Introducing `net/http` package
  - Overview of Go standard library
    - `fmt`
    - `io`
    - `os`
    - `net/http`
    - `encoding/json`
    - `template/html`

- `struct` tags
  - `json:"-"`

- Introducing `mux`, `negroni` & `gorilla` libraries for writing ReSTful APIs
  - Comparison with [Gin](https://github.com/gin-gonic/gin)

- [Sample App](https://github.com/algogrit/yaes-server)
  - config
  - logging
  - Error reporting
  - JWT token

- Working with Databases: `database/sql`
  - "ORM" in Go
    - `jinzhu/gorm` or `jmoiron/sqlx`

- [Hot Reloading](http://github.com/codegangsta/gin)
- [Awesome Go](https://awesome-go.com/)

- Deploying a Go app
  - Multi-stage docker builds
  - Deploy to heroku

- Alternative to ReST APIs
  - gRPC
    - [Protobuf](https://grpc.io/docs/guides/)
  - GraphQL

- Writing a [Websocket server using gorilla](http://www.gorillatoolkit.org/pkg/websocket)

- [HTTP/2](https://http2.github.io/)
  - [Server Push](https://blog.golang.org/h2push)

---

## References

- [gRPC Basics](https://grpc.io/docs/tutorials/basic/go/)
- [GraphQL](https://github.com/graphql-go/graphql)
- [Go in Cloud environments](https://rakyll.org/go-cloud/)
- [12 factor apps](https://12factor.net/)

## Assignment

- Writing a simple "Hello, World" API using native packages
- Writing a simple "Hello, World" API using either `gin` or `mux`
- Biblioteca
