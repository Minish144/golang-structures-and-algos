package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

func (n *Node) Traverse() {
	cur := n
	for cur != nil {
		fmt.Println("value: ", cur.Value)
		cur = cur.Next
	}
}

func (n *Node) Append(val int) {
	cur := n
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &Node{val, nil}
}

func (n *Node) Remove(val int) {
	dummyHead := &Node{Next: n}
	cur := dummyHead
	for cur != nil {
		if cur.Next != nil && cur.Next.Value == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	*n = *dummyHead.Next
}

func (n *Node) Lookup(val int) *Node {
	dummyHead := &Node{Next: n}
	cur := dummyHead
	for cur != nil {
		if cur.Next != nil && cur.Next.Value == val {
			return cur.Next
		} else {
			cur = cur.Next
		}
	}
	return nil
}

func main() {
	root := Node{5, nil}
	root.Append(10)
	root.Append(15)
	root.Append(20)
	root.Append(25)

	root.Traverse()

	fmt.Println(root.Lookup(10))
}
