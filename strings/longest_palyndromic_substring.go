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

// longestPalindrome brute force approach
func longestPalindrome(s string) string {
	n := len(s)
	var longestP string

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if j-i+1 > len(longestP) && isPalyndrome(s[i:j+1]) {
				longestP = s[i : j+1]
			}
		}
	}

	return longestP
}

func isPalyndrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i <= j && s[i] == s[j] {
		i += 1
		j -= 1
	}

	return j < i
}

func main() {
	fmt.Println(longestPalindrome("iopsajhffgvrnyitusobwcxgwlwniqchfnssqttdrnqqcsrigjsxkzcmuoiyxzerakhmexuyeuhjfobrmkoqdljrlojjjysfdslyvckxhuleagmxnzvikfitmkfhevfesnwltekstsueefbrddxrmxokpaxsenwlgytdaexgfwtneurhxvjvpsliepgvspdchmhggybwupiqaqlhjjrildjuewkdxbcpsbjtsevkppvgilrlspejqvzpfeorjmrbdppovvpzxcytscycgwsbnmspihzldjdgilnrlmhaswqaqbecmaocesnpqaotamwofyyfsbmxidowusogmylhlhxftnrmhtnnljjhhcfvywsqimqxqobfsageysonuoagmmviozeouutsiecitrmkypwknorjjiaasxfhsftypspwhvqovmwkjuehujofiabznpipidhfxpoustquzyfurkcgmioxacleqdxgrxbldcuxzgbcazgfismcgmgtjuwchymkzoiqhzaqrtiykdkydgvuaqkllbsactntexcybbjaxlfhyvbxieelstduqzfkoceqzgncvexklahxjnvtyqcjtbfanzgpdmucjlqpiolklmjxnscjcyiybdkgitxnuvtmoypcdldrvalxcxalpwumfx"))
}
