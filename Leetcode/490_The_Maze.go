package main

import "fmt"

func hasPath(maze [][]int, start []int, destination []int) bool {
	queue := make([][]int, 0)
	visited := make(map[[]int]bool)

	queue = append(queue, start)
	visited[start] = true

	for len(queue) > 0 {
		coordinate := queue[0]
		queue := queue[1:]

		// Add all possible four destinations to the queue in clockwise order

		// Right destination
		mostRightIndex := -1
		for i := coordinate[1] + 1; i < len(maze); i++ {
			if i != 0 {
				mostRightIndex = i
			}
		}

		if mostRightIndex != -1 {
			queue = append(queue, []int{coordinate[0], mostRightIndex})
		}

		// Left etc.
	}

	return false
}

func main() {
	maze := [][]int{
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1},
		{0, 0, 0, 0, 0},
	}

	//
	fmt.Println(hasPath(maze, []int{}, []int{}))
}
