package main

import "fmt"

func productExceptSelfHacky(nums []int) []int {
	result := make([]int, len(nums))
	n := len(nums)

	numberOfZeroes := 0
	totalProduct := 1

	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			totalProduct *= nums[i]
		} else {
			numberOfZeroes += 1
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			if numberOfZeroes == 1 {
				result[i] = totalProduct
			} else {
				result[i] = 0
			}

		} else {
			if numberOfZeroes > 0 {
				result[i] = 0
			} else {
				result[i] = totalProduct / nums[i]
			}
		}
	}

	return result
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{0, 1, 1}))
	fmt.Println(productExceptSelf([]int{0, 1, 0}))
}
