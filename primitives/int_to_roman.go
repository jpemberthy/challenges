package main

import (
	"bytes"
	"fmt"
	"math"
)

// inToRoman converts an integer to a roman number
func intToRoman(num int) string {
	var roman bytes.Buffer

	// break the integer into fractions, calculate and append the roman representation
	// of each fraction starting with the most significant one.
	digits := make([]int, 0)
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}

	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		fraction := digits[i] * int(math.Pow10(i))
		roman.WriteString(fractionToRoman(fraction))
	}

	return roman.String()
}

func fractionToRoman(fraction int) string {
	intToSymbol := map[int]string{
		0:    "",
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}
	if roman, found := intToSymbol[fraction]; found {
		return roman
	}

	var roman bytes.Buffer
	bases := []int{1, 5, 10, 50, 100, 500, 1000}

	base := 0
	i := 0
	n := len(bases)
	for i < n {
		if fraction < bases[i] {
			break
		}
		base = bases[i]
		i += 1
	}

	// when base is power of 10,
	//
	// we check if the fraction is between (base and base*3),
	// if it's then we sum up and keep appending the roman representation of the base until sum > fraction.
	// else, we fall into the substraction branch and just need to return this base and the next_base representation.
	powersOf10 := map[int]bool{1: true, 10: true, 100: true, 1000: true}
	sum := base
	if _, ok := powersOf10[base]; ok {
		if fraction <= base*3 {
			for sum <= fraction {
				roman.WriteString(intToSymbol[base])
				sum += base
			}
		} else {
			roman.WriteString(intToSymbol[base])
			roman.WriteString(intToSymbol[bases[i]])
		}
		return roman.String()
	}

	// when base is not power of 10,
	// we check if the fraction is between (base and [base + prev_base * 3]),
	// if it's, then we return the current base representation and keep appending the representation of the previous base until sum > fraction.
	// else, we fall into the substraction branch and just need to return the previous and next_base representations.
	prevBase := bases[i-2]
	if fraction <= (base + (bases[i-2])*3) {
		roman.WriteString(intToSymbol[base])
		for sum < fraction {
			roman.WriteString(intToSymbol[prevBase])
			sum += prevBase
		}
	} else {
		roman.WriteString(intToSymbol[prevBase])
		roman.WriteString(intToSymbol[bases[i]])
	}

	return roman.String()
}

func main() {
	for i := 1; i <= 3999; i++ {
		fmt.Printf("%d \t-->\t %s\n", i, intToRoman(i))
	}
}
