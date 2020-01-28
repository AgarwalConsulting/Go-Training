package main

import "fmt"

type Node struct {
	Value    int
	NextNode *Node
}

func main() {
	list := Node{}

	fmt.Println(list)
}
