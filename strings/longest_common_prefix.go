// Write a function to find the longest common prefix string amongst an array of strings.

package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	prefixes := make(map[string]int)

	// seed prefixes with first string. The longest common prefix
	// has to be one of these prefixes, or "".
	n := len(strs)
	if n > 0 {
		s := strs[0]
		m := len(s)
		for i := 0; i < m; i++ {
			prefixes[s[0:i+1]] += 1
		}
	}

	for i := 1; i < n; i++ {
		s := strs[i]
		m := len(s)
		hasCommonPrefix := false
		for j := 0; j < m; j++ {
			prefix := s[0 : j+1]

			// only add prefixes that are already in the map,
			// else break because the following prefixes of this
			// string are not common.
			if _, found := prefixes[prefix]; found {
				prefixes[prefix] += 1
				hasCommonPrefix = true
			} else {
				break
			}
		}

		if !hasCommonPrefix {
			return ""
		}
	}

	var longestPrefix string
	var longestPrefixLength int

	for prefix, count := range prefixes {
		if count == n && longestPrefixLength < len(prefix) {
			longestPrefix = prefix
			longestPrefixLength = len(prefix)
		}
	}

	return longestPrefix
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"abcd", "ab", "abd", "ab"}))
	fmt.Println(longestCommonPrefix([]string{"abcd", "ab", "abd", "ab"}))
	fmt.Println(longestCommonPrefix([]string{"abcd", "ab", "abd", "ab"}))
	fmt.Println(longestCommonPrefix([]string{"abcd", "ab", "abd", "ab"}))
	fmt.Println(longestCommonPrefix([]string{"", "b"}))
	fmt.Println(longestCommonPrefix([]string{"", "", "", "", "", "b"}))
}
