package main

import "fmt"

func binarySearch(nums []int, target int, left int, right int) int {
	n := right - left + 1

	if n == 0 {
		return -1
	}

	if nums[left] == target {
		return left
	}

	if nums[right] == target {
		return right
	}

	if n == 1 {
		return -1
	}

	if left >= right {
		return -1
	}

	mid := (right + left) / 2
	if nums[mid] == target {
		return mid
	}

	if nums[left] < target && nums[mid] > target {
		return binarySearch(nums, target, left, mid-1)
	}

	if nums[mid] < target && nums[right] > target {
		return binarySearch(nums, target, mid+1, right)
	}

	// if it's not in the sorted side, then attempt in the unsorted one.
	if nums[mid] > nums[right] {
		return binarySearch(nums, target, mid+1, right)
	}
	return binarySearch(nums, target, left, mid-1)
}

func search(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums)-1)
}

func main() {
	fmt.Println(search([]int{5, 1, 2, 3, 4}, 1))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 5))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println(search([]int{}, 3))
	fmt.Println(search([]int{-1, 2, 3, 4, 7, 10, 11}, 3))
	fmt.Println(search([]int{-1, 2, 3, 4, 7, 10, 11}, 10))
	fmt.Println(search([]int{-1, 2, 3, 4, 7, 10, 11}, 11))
}
