# Overview

## Building Services

### OSI Layers

![OSI Layers](assets/services/osi-layers.png)
> Image Credits: https://www.bmc.com/blogs/osi-model-7-layers/

### ReST

* Works well with HTTP 1.1
* Usually, is JSON based, protobuf is rarely used
* Supports Client-Server architecture well, some of the complex interactions become slower
  * For streaming people tend to use, websockets
* Generating client/server stubs requires support using third-party standards like `swagger` or even `openapi`
* Fitting some use-cases into a more ReSTful style can be tedious
* Well supported by the browser

#### Writing our first "Hello, World!" API

* `net/http`
  * `http.ListenAndServe("<HOSTNAME>:<PORT>", nil)`
  * `mux := http.NewServeMux()`
  * `http.HandleFunc(<pathString>, <handlerFunc>)`
* `func <handlerName>(w http.ResponseWriter, req *http.Request) {`
* `encoding/json`
  * `json.NewEncoder(w).Encode(<varToBeSerialized>)`
  * `json.NewDecoder(req.Body).Decode(&v)`

#### `gorilla/mux`

* `r := mux.NewRouter()`
* `r.HandleFunc(<path>, <handlerFunc>).Methods(<httpMethods>)`
* `http.ListenAndServe(":<PORT>", r)`
* Processing variables
  * `GET /people/{id}`
  * `mux.Vars(req)` returns `map[string]string`
    * Eg: `GET /people/1`
      * `mux.Vars(req)["id"] // "1"`
* `negroni` or `alice` for middleware

### gRPC

* Works on top of HTTP 2
* Usually, uses protobuf for messaging, JSON is optional
  * Uses *proto3* syntax
* Auto-generates client and server side stubs
* Supports both a unary (request-response) rpc as well as streaming (two-way communication)
* Middleware are known as interceptors
* Isn't well supported in the browser

#### Writing out first "Hello, World!" rpc

* Start by defining your service & messages in a proto file, with `syntax = "proto3"`
* Generate the client & server stubs, using `protoc`
* Implementing client
  * Connect to the server using address, `conn, err := grpc.Dial(address, ...opts)`
  * Create a new client from generated code, `pb.NewHelloServiceClient()`
* Implementing server
  * Create a tcp socket to listen on `lis, err := net.Listen("tcp", address)`
  * Create a new grpc server `s := grpc.NewServer()`
  * Register an implmentation of `pb.HelloServiceServer` using: `pb.RegisterHelloServiceServer(s, &HelloService{})`
  * Start listening on the tcp socket `s.Serve(lis)`

### HTTP 1.1 vs HTTP 2 Service (Layer 5, 6 & 7)

![HTTP 1.1 vs HTTP 2](assets/services/http_1_vs_2.png)
> Image Credits: https://blog.cloudflare.com/introducing-http2/

* Text-based protocol
* Human readable
* Useful for application to application communication

#### HTTP 2

* Secure by default
  * opportunistic encryption (*TLS for http:// URLs*)

#### HTTP 3

![HTTP 2 vs HTTP 3 Stack](assets/services/http2-stack-vs-http3-stack.png)
> Image Credits: https://kinsta.com/blog/http3/

### TCP vs UDP Service (Layer 4)

![TCP + TLS Handshake](assets/services/tcp_tls_handshake.png)
> Image Credits: https://blog.cloudflare.com/introducing-http2/

![TCP Server & Client Communication](assets/services/tcp-server-client-communication.gif)
> Image Credits: https://www.cse.iitk.ac.in/users/dheeraj/cs425/lec18.html

![TCP vs UDP](assets/services/tcp-vs-udp.svg)
> Image Credits: https://www.cloudflare.com/learning/ddos/glossary/user-datagram-protocol-udp/

* Ports are defined at this layer
* Allows for process to process communication

### IP (Layer 3)

![IPv6 vs IPv4](assets/services/ipv6-vs-ipv4.jpg)
> Image Credits: https://www.comparitech.com/blog/vpn-privacy/ipv6-vs-ipv4/

* Provides logical addressing
* Datagram sequencing

---

## Security

### [OWASP](https://owasp.org/)

#### [Top 10](https://owasp.org/www-project-top-ten/)

* Injection: *Injection flaws, such as SQL, NoSQL, OS, and LDAP injection, occur when untrusted data is sent to an interpreter as part of a command or query.*

![Bobby Tables](https://imgs.xkcd.com/comics/exploits_of_a_mom.png)

* Broken Authentication: *authentication and session management functions are often implemented incorrectly*

* Sensitive Data Exposure: *Sensitive data may be compromised without extra protection, such as encryption at rest or in transit, and requires special precautions when exchanged with the browser.*

* XML External Entities: *External entities can be used to disclose internal files using the file URI handler, internal file shares, internal port scanning, remote code execution, and denial of service attacks.*

* Broken Access Control: *Restrictions on what authenticated users are allowed to do are often not properly enforced.*

* Security Misconfiguration: *Security misconfiguration is the most commonly seen issue. This is commonly a result of insecure default configurations, incomplete or ad hoc configurations, open cloud storage, misconfigured HTTP headers, and verbose error messages containing sensitive information.*

* Cross-Site Scripting XSS: *XSS allows attackers to execute scripts in the victim’s browser which can hijack user sessions, deface web sites, or redirect the user to malicious sites.*

* Insecure Deserialization: *Insecure deserialization often leads to remote code execution.*

* Using Components with Known Vulnerabilities: *Components, such as libraries, frameworks, and other software modules, run with the same privileges as the application.*

* Insufficient Logging & Monitoring: *Most breach studies show time to detect a breach is over 200 days, typically detected by external parties rather than internal processes or monitoring.*

#### Best practices for Go (Basic Hardening)

* Validate input entries: *To validate user input, you can use native Go packages like `strconv` to handle string conversions to other data types. Go also has support for regular expressions with `regexp` for complex validations. Even though Go’s preference is to use native libraries, there are third-party packages like `validator`. With `validator`, you can include validations for structs or individual fields more easily.*

* Use HTML templates: *Go has the package html/template to encode what the app will return to the user.*

* Protect yourself from SQL injections: Use prepared statements

* Encrypt sensitive information: *there’s a Go package that includes robust implementations to encrypt information like `crypto`. OWASP has a few [recommendations](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html#leverage-an-adaptive-one-way-function) of which encryption algorithms to use, such as bcrypt, PDKDF2, Argon2, or scrypt.*

* Enforce HTTPS communication: Let's encrypt is the leading free certificate provider in the world. It has issued more than a [billion certificates](https://letsencrypt.org/2020/02/27/one-billion-certs.html), powering ~200 million sites across the globe. There is not a single reason to skip on TLS anymore.

* Be mindful with errors and logs: Extensive logging, being mindful of not logging sensitive information is important.

### Using Let's Encrypt

Let's encrypt uses the ACME ([Automatic certificate Management Environment](https://en.wikipedia.org/wiki/Automated_Certificate_Management_Environment)) protocol to issue SSL certs automatically. [`autocert`](https://godoc.org/golang.org/x/crypto/acme/autocert), a third-party package, comes in handy when configuring a go server with certificates.

* The standard port for HTTPS is 443
* You can run only HTTP, only HTTPS or both
* If the server doesn’t have a certificate, it’ll use HTTP API to ask Let’s Encrypt servers for it.
* You must set up DNS correctly.

### Authentication mechanisms

#### Basic Auth

* Sends credentials in plain text (TLS has it encrypted) over the headers
* Simple to implement, simple to bypass
* `req.BasicAuth()`
* `grpc_auth.AuthFromMD(ctx, "basic")`

#### [JWT](https://jwt.io/) - JSON Web Token

* *JSON Web Tokens are an open, industry standard RFC 7519 method for representing claims securely between two parties.*
* *it's a signed JSON object that does something useful (for example, authentication). It's commonly used for Bearer tokens in Oauth 2. A token is made of three parts, separated by .'s. The first two parts are JSON objects, that have been base64url encoded. The last part is the signature, encoded the same way.*
* They replace the now-obsolete method of maintaining client session info on the server side
* JWT token can contain user info within them
