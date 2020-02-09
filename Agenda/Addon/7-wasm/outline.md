# WASM - Outline

- Introduction to [WebAssembly](https://webassembly.org/)

- Writing your [own wasm application in Go](https://github.com/golang/go/wiki/WebAssembly)
  - package [`syscall/js`](https://golang.org/pkg/syscall/js/)
  - `js.Value`
  - `js.Global()`
    - `.Get`
    - `.Set`
    - `.Call`
  - `js.FuncOf`
  - Compiling
    - `GOOS=js` & `GOARCH=wasm`

- Looking at a [sample application](https://github.com/Chennai-Golang/wasm-go-sample)

---

## References

- [WebAssembly in Go](https://github.com/golang/go/wiki/WebAssembly)
