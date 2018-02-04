// There are two sorted arrays nums1 and nums2 of size m and n respectively.

// Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
package main

import "fmt"

// findMedianSortedArrays O(n)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var i, j, k int
	nums3 := make([]int, len(nums1)+len(nums2))

	for i < len(nums1) || j < len(nums2) {
		if i == len(nums1) {
			nums3[k] = nums2[j]
			j += 1
		} else if j == len(nums2) {
			nums3[k] = nums1[i]
			i += 1
		} else if nums1[i] <= nums2[j] {
			nums3[k] = nums1[i]
			i += 1
		} else {
			nums3[k] = nums2[j]
			j += 1
		}

		k += 1
	}

	// If there is an even number of observations, then there is no single middle value;
	// the median is then usually defined to be the mean of the two middle values.
	n := len(nums3)
	middle := n / 2
	var median float64
	if n%2 == 0 {
		median = float64(nums3[middle-1]+nums3[middle]) / 2.0
	} else {
		median = float64(nums3[middle])
	}

	return median
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
