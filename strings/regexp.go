// '.' Matches any single character.
// '*' Matches zero or more of the preceding element.

// The matching should cover the entire input string (not partial).

// The function prototype should be:
// bool isMatch(const char *s, const char *p)

// Some examples:
// isMatch("aa","a") → false
// isMatch("aa","aa") → true
// isMatch("aaa","aa") → false
// isMatch("aa", "a*") → true
// isMatch("aa", ".*") → true
// isMatch("ab", ".*") → true
// isMatch("aab", "c*a*b") → true
package main

import "fmt"

const zeroOrMoreOfPreviousToken = "*"
const anyToken = "."

func isMatch(s string, p string) bool {
	lenP := len(p)
	lenS := len(s)
	if lenP == 0 {
		return lenS == 0
	}

	firstMatch := lenS > 0 && (s[0] == p[0] || string(p[0]) == anyToken)
	if lenP >= 2 && string(p[1]) == zeroOrMoreOfPreviousToken {
		return isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
	} else {
		return firstMatch && isMatch(s[1:], p[1:])
	}
}

func main() {
	fmt.Println(isMatch("aa", "a") == false)
	fmt.Println(isMatch("aa", "aa") == true)
	fmt.Println(isMatch("aaa", "aa") == false)
	fmt.Println(isMatch("aa", "a*") == true)
	fmt.Println(isMatch("aa", ".*") == true)
	fmt.Println(isMatch("ab", ".*") == true)
	fmt.Println(isMatch("aab", "c*a*b") == true)
}
