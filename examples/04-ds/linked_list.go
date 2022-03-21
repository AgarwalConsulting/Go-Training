package main

import "fmt"

type Node struct {
	Value    string
	NextNode *Node
}

func (list *Node) Insert(val string) {
	list.NextNode = &Node{Value: val, NextNode: list.NextNode}
}

func main() {
	list := Node{Value: "Gaurav", NextNode: nil}

	list.Insert("Rajat")
	list.Insert("Arpit")

	fmt.Println(list)
}
