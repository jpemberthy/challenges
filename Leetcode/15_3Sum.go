// https://leetcode.com/problems/3sum/description/
// Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

// Note: The solution set must not contain duplicate triplets.

// For example, given array S = [-1, 0, 1, 2, -1, -4]

// A solution set is:
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]

package main

import (
	"fmt"
	"sort"
)

func threeSumN3(nums []int) [][]int {
	set := make(map[string]bool)
	n := len(nums)
	iLimit := n - 2
	jLimit := n - 1

	solution := make([][]int, 0)

	for i := 0; i < iLimit; i++ {
		for j := i + 1; j < jLimit; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					ary := []int{nums[i], nums[j], nums[k]}
					sort.Ints(ary)
					key := fmt.Sprintf("%d-%d-%d", ary[0], ary[1], ary[2])
					if !set[key] {
						solution = append(solution, ary)
						set[key] = true
					}
				}
			}
		}
	}

	return solution
}

func main() {
	solutionN3 := threeSumN3([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println("Solution: ", solutionN3)
}
