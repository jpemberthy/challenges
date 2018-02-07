// Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

// Example:

// Input: "babad"

// Output: "bab"

// Note: "aba" is also a valid answer.

// Example:

// Input: "cbbd"

// Output: "bb"
package main

import "fmt"

var cache map[string]bool

// longestPalindrome brute force with index-based caching
func longestPalindrome(s string) string {
	cache = make(map[string]bool)

	n := len(s)
	var longestP string

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if j-i+1 > len(longestP) && isPalyndrome(s, i, j) {
				longestP = s[i : j+1]
			}
		}
	}

	return longestP
}

func isPalyndrome(s string, i int, j int) bool {
	key := fmt.Sprintf("%d-%d", i, j)

	if _isPalyndrome, found := cache[key]; found {
		return _isPalyndrome
	}

	n := j - i + 1
	if n == 1 {
		cache[key] = true
		return true
	}

	if n == 2 && s[i] == s[j] {
		cache[key] = true
		return true
	}

	var _isPalyndrome bool
	if s[i] == s[j] {
		_isPalyndrome = isPalyndrome(s, i+1, j-1)
	}

	cache[key] = _isPalyndrome
	return _isPalyndrome
}

func main() {
	fmt.Println(longestPalindrome("foo"))
}
