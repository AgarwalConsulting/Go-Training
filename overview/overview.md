# Overview

## Sweets spots of Go

* Why Go?

### Compiling cross-platform binaries

* `GOOS` & `GOARCH` environment variables

### Concurrency vs Parallelism

OS Overview

![OS Overview](assets/1-OS-process-and-its-threads.png)

> Image Credits: https://github.com/sathishvj/optimizing-go-programs#m-p-g

Concurrency vs Parallelism

![Concurrency vs Parallelism](assets/4-concurrency-and-parallelism.png)

> Image Credits: https://github.com/sathishvj/optimizing-go-programs#m-p-g

* Overview of Go scheduler

OS vs Go Scheduler

![OS vs Go Scheduler](assets/go-vs-os-scheduler.png)

> Image Credits: https://speakerdeck.com/kavya719/the-scheduler-saga?slide=15

OS Thread & Go Routine Scheduling

![Goroutine vs OS Thread](assets/os-thread-goroutine.png)
> Image Credits: https://speakerdeck.com/kavya719/the-scheduler-saga?slide=18

* `GOMAXPROCS` controls the max number of OS threads a go program can create

### Other Packages

* `context` package
  * Useful for signaling multiple goroutines at the same time
  * Useful for controlling goroutine execution
  * Useful for sharing values across goroutine boundaries

---

## Building Services

### OSI Layers

![OSI Layers](assets/services/osi-layers.png)
> Image Credits: https://www.bmc.com/blogs/osi-model-7-layers/

### ReST

* Works well with HTTP 1.1
* Usually, is JSON based, protobuf is rarely used

### gRPC

* Works on top of HTTP 2
* Usually, uses protobuf for messaging, JSON is optional
  * Uses *proto3* syntax
* Auto-generates client and server side stubs

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

#### [JWT](https://jwt.io/) - JSON Web Token

* *JSON Web Tokens are an open, industry standard RFC 7519 method for representing claims securely between two parties.*
* They replace the now-obsolete method of maintaining client session info on the server side
* JWT token can contain user info within them
