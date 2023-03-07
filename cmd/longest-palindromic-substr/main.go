package main

import "fmt"

/*
Longest Palindromic Substring
https://leetcode.com/problems/longest-palindromic-substring/

Алгоритм взят отсюда https://www.geeksforgeeks.org/given-a-string-print-all-possible-palindromic-partition/
Решение не проходит по таймингу ;)
*/

var longest string

func main() {
	test := "ABCDCBEFAFEBC"
	test = "abacab"

	result := longestPalindrome(test)
	fmt.Println(result)
}

func longestPalindrome(s string) string {
	longest = ""
	recursive([]rune(s), 0, len(s))

	return longest
}

func recursive(str []rune, start, n int) {
	if start >= n {
		return
	}

	for i := start; i < n; i++ {
		if isPalindrome(str, start, i) {
			if (i - start + 1) > len(longest) {
				var palindrome string
				for j := start; j <= i; j++ {
					palindrome += string(str[j])
				}

				longest = palindrome
			}

			recursive(str, i+1, n)
		}
	}

}

func isPalindrome(str []rune, low, high int) bool {
	for low < high {
		if str[low] != str[high] {
			return false
		}
		low++
		high--
	}

	return true
}
