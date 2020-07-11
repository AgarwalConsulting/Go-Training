package generator

import "sync"

type FibonacciGenerator struct {
	i, j int64
	mut  *sync.Mutex
}

func (g *FibonacciGenerator) NextValue() int64 {
	g.mut.Lock()
	defer g.mut.Unlock()
	g.i, g.j = g.j, g.i+g.j

	return g.j - g.i
}

func (g *FibonacciGenerator) Reset() {
	g.mut.Lock()
	defer g.mut.Unlock()
	g.i, g.j = 0, 1
}

func NewFibonacciGenerator() *FibonacciGenerator {
	return &FibonacciGenerator{0, 1, &sync.Mutex{}}
}
