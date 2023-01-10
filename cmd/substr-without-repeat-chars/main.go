package main

import (
	"fmt"
)

/*
Task #3. Longest Substring Without Repeating Characters. Medium level.
https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
Runtime - 3ms(beats 94%), memory: 3.1mb(beats 55%)
*/
func main() {
	s := "aab"
	result := lengthOfLongestSubstring(s)
	fmt.Println(result)
}

func lengthOfLongestSubstring(s string) int {
	var bigSubstr, j, idx int
	substr := make(map[rune]int)
	ok := false
	compare := func(first, sec int) int {
		if first > sec {
			return first
		}
		return sec
	}

	for i, char := range s {
		if j, ok = substr[char]; ok && j >= idx {
			bigSubstr = compare(i-idx, bigSubstr)
			idx = j + 1
		}

		substr[char] = i
	}

	return compare(len(s)-idx, bigSubstr)
}
