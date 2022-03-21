package main

import "fmt"

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

func (self *Node[T]) Insert(val T) *Node[T] {
	newHead := Node[T]{Value: val, Next: self}

	return &newHead
}

func (curr *Node[T]) Display() {
	if curr == nil {
		return
	}
	fmt.Println("\t", curr.Value)
	curr.Next.Display()
}

func NewLinkedList[T any]() *Node[T] {
	return nil
}

func main() {
	list := NewLinkedList[string]()

	list = list.Insert("Gaurav")

	fmt.Println("-----------------------------")
	list.Display()

	list = list.Insert("Sid")

	fmt.Println("-----------------------------")
	list.Display()

	list = list.Insert("Ashish")

	fmt.Println("-----------------------------")
	list.Display()

	fmt.Println("-----------------------------")
}
