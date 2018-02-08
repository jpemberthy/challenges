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

// longestPalindrome brute force cached approach with dynamic programming
func longestPalindrome(s string) string {
	n := len(s)
	palindromes := make([][]bool, n)
	for i := range palindromes {
		palindromes[i] = make([]bool, n)
	}

	var max, left, right int

	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if s[j] == s[i] && (i-j+1 <= 2 || palindromes[j+1][i-1]) {
				palindromes[j][i] = true
			}

			if palindromes[j][i] && i-j+1 > max {
				max = i - j + 1
				left = j
				right = i
			}
		}
	}

	return s[left : right+1]
}

func main() {
	fmt.Println(longestPalindrome("iopsajhffgvrnyitusobwcxgwlwniqchfnssqttdrnqqcsrigjsxkzcmuoiyxzerakhmexuyeuhjfobrmkoqdljrlojjjysfdslyvckxhuleagmxnzvikfitmkfhevfesnwltekstsueefbrddxrmxokpaxsenwlgytdaexgfwtneurhxvjvpsliepgvspdchmhggybwupiqaqlhjjrildjuewkdxbcpsbjtsevkppvgilrlspejqvzpfeorjmrbdppovvpzxcytscycgwsbnmspihzldjdgilnrlmhaswqaqbecmaocesnpqaotamwofyyfsbmxidowusogmylhlhxftnrmhtnnljjhhcfvywsqimqxqobfsageysonuoagmmviozeouutsiecitrmkypwknorjjiaasxfhsftypspwhvqovmwkjuehujofiabznpipidhfxpoustquzyfurkcgmioxacleqdxgrxbldcuxzgbcazgfismcgmgtjuwchymkzoiqhzaqrtiykdkydgvuaqkllbsactntexcybbjaxlfhyvbxieelstduqzfkoceqzgncvexklahxjnvtyqcjtbfanzgpdmucjlqpiolklmjxnscjcyiybdkgitxnuvtmoypcdldrvalxcxalpwumfx"))
}
