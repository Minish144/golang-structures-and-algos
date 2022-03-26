package main

import "fmt"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func (t *TreeNode) Add(val int) {
	if t.value == val {
		return
	} else if val < t.value {
		if t.left == nil {
			t.left = &TreeNode{value: val}
		} else {
			t.left.Add(val)
		}
	} else {
		if t.right == nil {
			t.right = &TreeNode{value: val}
		} else {
			t.right.Add(val)
		}
	}
}

func (t *TreeNode) Traverse() {
	if t == nil {
		return
	}
	t.left.Traverse()
	fmt.Println(t.value)
	t.right.Traverse()
}

func (t *TreeNode) Lookup(val int) *TreeNode {
	if t == nil {
		return nil
	}

	if t.value == val {
		return t
	} else if val < t.value {
		return t.left.Lookup(val)
	} else {
		return t.right.Lookup(val)
	}
}

func main() {
	tr := TreeNode{50, nil, nil}
	tr.Add(10)
	tr.Add(30)
	tr.Add(20)
	tr.Add(70)
	tr.Add(80)
	tr.Add(90)
	tr.Add(75)

	tr.Traverse()
	fmt.Printf("%+v", tr.Lookup(70))
}
