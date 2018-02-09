package main // Given a 32-bit signed integer, reverse digits of an integer.
import (
	"fmt"
	"math"
)

// Example 1:

// Input: 123
// Output:  321
// Example 2:

// Input: -123
// Output: -321
// Example 3:

// Input: 120
// Output: 21
// Note:
// Assume we are dealing with an environment which could only hold integers within the 32-bit signed integer range. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

func reverse(x int) int {
	var negative bool
	if x < 0 {
		negative = true
		x *= -1
	}

	// the max integer allowed is +-2^31, the remaining bit is used
	// for the sign.
	max := int(math.Pow(2, 31))
	if x > max {
		return 0
	}

	digits := make([]int, 0)
	for x > 0 {
		digits = append(digits, x%10)
		x = x / 10
	}

	var r int
	n := len(digits)
	for i, digit := range digits {
		r += digit * int(math.Pow10(n-i-1))
	}

	if r > max {
		return 0
	}

	if negative {
		r *= -1
	}

	return r
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(1534236469))
	fmt.Println(reverse(-2147483648))
}
