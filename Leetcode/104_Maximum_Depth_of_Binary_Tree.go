package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	_maxDepth := -1

	stack := make([]*TreeNode, 0)
	visited := make(map[*TreeNode]bool)

	stack = append(stack, root)
	visited[root] = true

	for len(stack) > 0 {
		node := stack[len(stack)-1]

		if node.Left == nil && node.Right == nil {
			if len(stack) > _maxDepth {
				_maxDepth = len(stack)
			}
			stack = stack[0 : len(stack)-1]
			continue
		}

		if node.Left != nil && !visited[node.Left] {
			stack = append(stack, node.Left)
			visited[node.Left] = true
			continue
		}

		if node.Right != nil && !visited[node.Right] {
			stack = append(stack, node.Right)
			visited[node.Right] = true
			continue
		}

		stack = stack[0 : len(stack)-1]
	}

	return _maxDepth
}

func main() {
	root := &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}

	fmt.Println(maxDepth(root))
}
