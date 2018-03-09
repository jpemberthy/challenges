// Given a roman numeral, convert it to an integer.
// Input is guaranteed to be within the range from 1 to 3999.
package main

import "fmt"

func romanToInt(s string) int {
	symbolToI := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	n := len(s)
	number := 0
	i := 0
	var currentBase int

	for i < n {
		currentSum := 0
		prevBase := symbolToI[string(s[i])]
		for j := i; j < n; j++ {
			currentBase = symbolToI[string(s[j])]
			if currentBase != prevBase {
				break
			}
			i += 1
			currentSum += currentBase
		}

		// e.g IV or IX
		if prevBase < currentBase {
			currentSum = currentBase - prevBase
			i += 1
		}

		number += currentSum
	}

	return number
}

func main() {
	fmt.Println(romanToInt("XXVII"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("XLIX"))
}
