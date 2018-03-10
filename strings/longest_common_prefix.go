// Write a function to find the longest common prefix string amongst an array of strings.

package main

import "fmt"

// "abcd" --> "a" --> "ab" --> "abc", "abcd"
// "ab" --> "a" --> "ab"
// "ad" --> "a" --> "ad"
// "a" --> "c"
func longestCommonPrefix(strs []string) string {
	prefixes := make(map[string]int)

	for _, s := range strs {
		n := len(s)
		for i := 0; i < n; i++ {
			prefix := s[0 : i+1]
			prefixes[prefix] += 1
		}
	}

	var longestPrefix string
	var longestPrefixLength int
	n := len(strs)

	for prefix, count := range prefixes {
		if count == n && longestPrefixLength < len(prefix) {
			longestPrefix = prefix
			longestPrefixLength = len(prefix)
		}
	}

	return longestPrefix
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"abcd", "ab", "ad", "a"}))
	fmt.Println(longestCommonPrefix([]string{"", "b"}))
}
