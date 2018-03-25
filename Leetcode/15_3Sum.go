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

func threeSum(nums []int) [][]int {
	// sorting the input is an optimization to help with the runtime constraint in Leetcode.
	sort.Ints(nums)
	solution := make([][]int, 0)
	triplets := make(map[string]bool)
	// complements is map of the form { number --> indexes_in_nums }
	// it's a reverse to lookup if and where the zero complements of a tuple are present.
	complements := make(map[int][]int)
	for index, value := range nums {
		complements[value] = append(complements[value], index)
	}

	n := len(nums)
	iLimit := n - 1
	for i := 0; i < iLimit; i++ {
		if i+2 < n && nums[i] == nums[i+2] && n-i > 3 {
			continue
		}
		for j := i + 1; j < n; j++ {
			c := (nums[i] + nums[j]) * -1
			if complementIndexes, found := complements[c]; found {
				for _, k := range complementIndexes {
					if k != i && k != j {
						ary := []int{nums[i], nums[j], nums[k]}
						sort.Ints(ary)
						key := fmt.Sprintf("%d-%d-%d", ary[0], ary[1], ary[2])
						if !triplets[key] {
							solution = append(solution, ary)
							triplets[key] = true
						}
						break
					}
				}
			}
		}
	}

	return solution
}

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
	solution := threeSum([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	fmt.Println("Solution: ", solution)
}
