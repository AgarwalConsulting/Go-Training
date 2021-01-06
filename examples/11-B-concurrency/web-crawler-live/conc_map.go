package main

import "sync"

type ConcurrentMap struct {
	m   map[string]bool
	mut *sync.RWMutex
}

// IsPresent
// Set

// Value Receiver => Pointer Receiver
func (cm *ConcurrentMap) IsPresent(key string) bool {
	cm.mut.RLock()
	isPresent := cm.m[key]
	cm.mut.RUnlock()
	return isPresent
}

func (cm *ConcurrentMap) Set(key string) {
	cm.mut.Lock()
	cm.m[key] = true
	cm.mut.Unlock()
}

func (cm *ConcurrentMap) CheckAndSet(key string) bool {
	cm.mut.Lock()
	defer cm.mut.Unlock()

	isPresent := cm.m[key]
	cm.m[key] = true

	return isPresent
}

func NewConcurrentMap() ConcurrentMap {
	return ConcurrentMap{m: make(map[string]bool), mut: new(sync.RWMutex)}
}
