package main // Given a string, find the length of the longest substring without repeating characters.
import "fmt"

// Examples:

// Given "abcabcbb", the answer is "abc", which the length is 3.

// Given "bbbbb", the answer is "b", with the length of 1.

// Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
func lengthOfLongestSubstring(s string) int {
	cache := make(map[byte]bool)
	var length, i, j int
	n := len(s)

	for i < n && j < n {
		if cache[s[j]] {
			delete(cache, s[i])
			i = i + 1
		} else {
			cache[s[j]] = true
			j = j + 1
			if length < j-i {
				length = j - i
			}
		}
	}

	return length
}

func main() {
	fmt.Println(lengthOfLongestSubstring("c") == 1)
	fmt.Println(lengthOfLongestSubstring("abcabcbb") == 3)
	fmt.Println(lengthOfLongestSubstring("bbbbb") == 1)
	fmt.Println(lengthOfLongestSubstring("pwwkew") == 3)
}
