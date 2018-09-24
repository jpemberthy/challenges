package main

import (
	"fmt"
	"math"
)

func myPow(x float64, n int) float64 {
	if x == -1 {
		return -1
	}

	if x == 1 || n == 0 {
		return 1
	}

	if n < 0 {
		return 1.0 / myPow(x, -n)
	}

	result := x
	for i := 1; i < n; i++ {
		result *= x

		if result > math.MaxInt32 {
			result = math.Inf(1)
			break
		}
	}

	if x < 0 && result > 0 && n%2 != 0 {
		return -result
	}

	return result
}

func main() {
	fmt.Println(myPow(2, -2147483648))
	fmt.Println(myPow(2, 10))
}
