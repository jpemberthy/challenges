package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	n := len(height)
	var area int

	i := 0
	j := n - 1
	prevH := 0
	for i < j {
		leftH := height[i]
		rightH := height[j]
		minH := min(leftH, rightH)
		if minH > prevH {
			_area := minH * (j - i)
			if _area > area {
				area = _area
			}
		}

		if minH == height[i] {
			i += 1
		} else {
			j -= 1
		}
		prevH = minH
	}

	return area
}

func main() {
	fmt.Println(maxArea([]int{3, 4, 2, 4, 5, 1, 1, 1}))
}
