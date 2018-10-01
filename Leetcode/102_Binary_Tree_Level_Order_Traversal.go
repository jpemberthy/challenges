package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		levelNodes := make([]int, 0)
		nodesInCurrentLevel := len(queue)

		for i := 0; i < nodesInCurrentLevel; i++ {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			queue = queue[1:]
			levelNodes = append(levelNodes, node.Val)
		}

		result = append(result, levelNodes)
	}

	return result
}

func main() {
	root := &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}

	fmt.Println(levelOrder(root))
}
