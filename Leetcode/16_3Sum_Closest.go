// Given an array S of n integers, find three integers in S such that the sum is closest to a given number, target. Return the sum of the three integers. You may assume that each input would have exactly one solution.

//     For example, given array S = {-1 2 1 -4}, and target = 1.

//     The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
package main

import (
	"fmt"
	"math"
)

func threeSumClosest(nums []int, target int) int {
	var sum int
	delta := math.MaxInt32
	n := len(nums)
	iLimit := n - 2
	jLimit := n - 1
	for i := 0; i < iLimit; i++ {
		for j := i + 1; j < jLimit; j++ {
			for k := j + 1; k < n; k++ {
				tripletSum := nums[i] + nums[j] + nums[k]
				if tripletSum == target {
					return tripletSum
				}

				tripletDelta := tripletSum - target
				if tripletDelta < 0 {
					tripletDelta *= -1
				}
				if tripletDelta < delta {
					delta = tripletDelta
					sum = tripletSum
				}
			}
		}
	}

	return sum
}

func main() {
	fmt.Println(threeSumClosest([]int{1, 1, 1, 0}, -100))
	fmt.Println(threeSumClosest([]int{1, 1, 1, 0}, 100))
}
