// Implement atoi to convert a string to an integer.
// Hint: Carefully consider all possible input cases. If you want a challenge, please do not see below and ask yourself what are the possible input cases.

// Notes: It is intended for this problem to be specified vaguely (ie, no given input specs). You are responsible to gather all the input requirements up front.

// The requirements below are to make the program pass on leetcode. That said, the function could be simplified if we were to change the
// signature of the function to return an err on invalid inputs or size. Just like the strconv.Atoi function of go.

// Requirements for atoi:

// The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.

// The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.

// If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.

// If no valid conversion could be performed, a zero value is returned. If the correct value is out of the range of representable values, INT_MAX (2147483647) or INT_MIN (-2147483648) is returned.

package main

import (
	"fmt"
	"math"
	"unicode"
)

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}

	digits := map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}

	// first sanitization for leetcode requirements.
	// remove all whitespaces
	n := len(str)
	for i, char := range str {
		if !unicode.IsSpace(char) {
			str = str[i:n]
			break
		}
	}

	signOrFirstDigit := string(str[0])
	var signed bool

	if signOrFirstDigit == "+" || signOrFirstDigit == "-" {
		signed = true
	}

	signConversor := 1
	if signed && signOrFirstDigit == "-" {
		signConversor = -1
	}

	// if signed, shift the string.
	if signed {
		str = str[1:len(str)]
	}

	parsedDigits := make([]int, 0)
	for _, digit := range str {
		if intDigit, found := digits[string(digit)]; found {
			parsedDigits = append(parsedDigits, intDigit)
		} else {
			// we're done processing the number or we have found a
			// non digit char.
			break
		}
	}

	intMax := 2147483647
	intMin := -2147483648

	var output int
	n = len(parsedDigits)

	for i, digit := range parsedDigits {
		power := n - i - 1
		if digit != 0 && power >= 10 {
			if signConversor == 1 {
				return intMax
			}
			return intMin
		}

		output += digit * int(math.Pow10(power))
		if output > intMax && signConversor == 1 {
			return intMax
		} else if output >= intMax+1 && signConversor == -1 {
			return intMin
		}
	}

	return output * signConversor
}

func main() {
	fmt.Println(myAtoi("18446744073709551617"))
	fmt.Println(myAtoi("-10"))
	fmt.Println(myAtoi("324442"))
	fmt.Println(myAtoi("-2147483650"))
	fmt.Println(myAtoi("-2147483647"))
}
