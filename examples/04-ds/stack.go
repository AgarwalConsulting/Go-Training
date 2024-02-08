package main

import (
	"errors"
	"fmt"
)

type Stack[T any] []T

func (s *Stack[T]) Push(el ...T) {
	*s = append(*s, el...)
}

func (s *Stack[T]) Pop() (T, error) {
	if len(*s) == 0 {
		return *new(T), errors.New("last element")
	}
	lastEl := (*s)[len(*s)-1]

	*s = (*s)[0 : len(*s)-1]

	return lastEl, nil
}

func NewStack[T any]() Stack[T] {
	return make(Stack[T], 0, 10)
}

func main() {
	s := NewStack[int]()

	s.Push(1, 2, 3, 5, 7)

	for {
		el, err := s.Pop()
		if err != nil {
			break
		}

		fmt.Println("Popped:", el)
	}
}
