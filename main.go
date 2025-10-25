package main

import "fmt"

type Node struct {
	data	int16
	next	*Node
}

var head *Node = nil
var curr *Node

func AddNode(data int16) {
	newNode := &Node{data, nil}

	if head == nil {
		head = newNode
	} else {
		curr.next = newNode
	}
	curr = newNode
}

func Traverse() {
	ptr := head
	for ptr != nil {
		fmt.Printf("%d ", ptr.data)
		ptr = ptr.next
	}
	fmt.Printf("\n")
}

func main() {
	fmt.Printf("Hello, world\n")

	AddNode(5)
	AddNode(6)
	AddNode(7)
	AddNode(8)
	AddNode(9)

	Traverse()
}
