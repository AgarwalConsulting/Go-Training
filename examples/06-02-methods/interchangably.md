# Receiver functions can be called interchangably

value receiver: `(v Vertex) Abs`
pointer receiver: `(v *Vertex) Scale`

value: `var vertex Vertex`
pointer: `var p *Vertex`

```golang
vertex.Abs()
vertex.Scale()

p.Abs()
p.Scale()
```
