package main

import (
	"fmt"
	"math"
)

type Node struct {
	index int
	edges []*Node
}

type Graph map[int]*Node

func buildGraph(nums []int) Graph {
	graph := make(Graph)
	n := len(nums)
	for i := 0; i < n; i++ {
		node := graph[i]
		if node == nil {
			node = &Node{index: i, edges: make([]*Node, 0)}
		}
		graph[i] = node

		for j := 0; j < nums[i] && j+i+1 < n; j++ {
			k := i + j + 1
			connection := graph[k]
			if connection == nil {
				connection = &Node{index: k, edges: make([]*Node, 0)}
				graph[k] = connection
			}

			node.edges = append(node.edges, connection)
		}
	}

	return graph
}

func bellmanFord(graph Graph) []int {
	n := len(graph)
	distances := make([]int, n)
	for i := 0; i < n; i++ {
		distances[i] = math.MaxInt32
	}

	distances[0] = 0

	for {
		updated := false
		for j := 0; j < n; j++ {
			node := graph[j]
			for _, edgeNode := range node.edges {
				// in this case the weight is always 1
				distance := distances[j] + 1
				if distance < distances[edgeNode.index] {
					distances[edgeNode.index] = distance
					updated = true
				}
			}
		}

		if !updated {
			break
		}
	}

	return distances
}

func jump(nums []int) int {
	graph := buildGraph(nums)
	d := bellmanFord(graph)
	return d[len(graph)-1]
}

func main() {
	fmt.Println(jump([]int{1, 3, 1, 1, 4}))
	fmt.Println("finished")
}
