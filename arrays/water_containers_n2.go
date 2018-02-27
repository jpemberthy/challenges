// Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

// Note: You may not slant the container and n is at least 2.
package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// maxArea n^2 solution.
func maxArea(height []int) int {
	var max int
	n := len(height)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			minHeight := min(height[i], height[j])
			area := minHeight * (j - i)
			if area > max {
				max = area
			}
		}
	}

	return max
}

func main() {
	fmt.Println(maxArea([]int{4, 3, 6, 2, 7, 5}))
}
