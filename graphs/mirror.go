package main

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// IsMirror checks if a BTree is mirror image or symmetric.
func IsMirror(n1 *Node, n2 *Node) bool {
	if n1 == nil && n2 == nil {
		return true
	}

	if n1.Left == nil && n2.Right != nil {
		return false
	}

	if n1.Right == nil && n2.Left != nil {
		return false
	}

	if n2.Right == nil && n1.Left != nil {
		return false
	}

	if n2.Left == nil && n1.Right != nil {
		return false
	}

	if n1.Value == n2.Value {
		return IsMirror(n1.Left, n2.Right) && IsMirror(n1.Right, n2.Left)
	}

	return false
}

func main() {
	root := Node{Value: 1}

	root.Left = &Node{
		Value: 2,
		Left:  &Node{Value: 3, Left: &Node{Value: 5}},
		Right: &Node{Value: 4},
	}

	root.Right = &Node{
		Value: 2,
		Left:  &Node{Value: 4},
		Right: &Node{Value: 3, Right: &Node{Value: 5}},
	}

	fmt.Printf("Mirror? %v", IsMirror(root.Left, root.Right))
}
