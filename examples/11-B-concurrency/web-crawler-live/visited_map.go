package main

import "sync"

type VisitedMap struct {
	m   map[string]bool
	mut sync.Mutex
}

func (vm *VisitedMap) IsVisitedAndSet(url string) bool {
	vm.mut.Lock()
	defer vm.mut.Unlock()

	v := vm.m[url]
	vm.m[url] = true
	return v
}

func NewVisitedMap() VisitedMap {
	return VisitedMap{m: make(map[string]bool)}
}
