// Determine whether an integer is a palindrome. Do this without extra space.

package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	// negative int are not palindrome
	if x < 0 {
		return false
	}

	if x/10 == 0 {
		return true
	}

	_x := x
	var reverse int
	largestPower := int(math.Pow10(int(math.Log10(float64(x)))))
	if x%10 == x/largestPower {
		for _x > 0 {
			reverse += (_x % 10) * largestPower
			largestPower /= 10
			_x /= 10
		}
	}

	return x == reverse
}

func main() {
	fmt.Println(isPalindrome(10))
}
