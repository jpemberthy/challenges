// Given a string, find the length of the longest substring without repeating characters.
// Examples:

// Given "abcabcbb", the answer is "abc", which the length is 3.

// Given "bbbbb", the answer is "b", with the length of 1.

// Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
package main

import "fmt"

var cache map[string]bool

func hasRepeatedChars(s string) bool {
	if cache[s] {
		return true
	}

	chars := make(map[rune]bool)
	for _, char := range s {
		if chars[char] {
			cache[s] = true
			return true
		}
		chars[char] = true
	}

	return false
}

func lengthOfLongestSubstring(s string) int {
	// Kinda hacky init but works for leetcode.
	if cache == nil {
		cache = make(map[string]bool)
	}

	length := 0
	for i := 0; i < len(s) || len(s)-i > length; i++ {
		for j := len(s); j > i; j-- {
			_length := j - i
			if _length > length {
				substring := s[i:j]
				if hasRepeatedChars(substring) {
					cache[substring] = true
					continue
				}

				if _length > length {
					length = _length
				}
			} else {
				break
			}
		}
	}

	return length
}

func main() {
	fmt.Println(lengthOfLongestSubstring("aaaaa") == 1)
	fmt.Println(lengthOfLongestSubstring("bbbbb") == 1)
	fmt.Println(lengthOfLongestSubstring("pwwkew") == 3)
}
