package main

import (
	"fmt"
)

func main() {
	s := "catsandog"
	wordDict := []string{"cats","dog","sand","and","cat", "an"}

	//s = "applepenapple"
	//wordDict = []string{"apple","pen", "app", "a", "epen", "pple", "appl"}

	//s = "ааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааа"
	//wordDict = []string{"а","аа","ааа","аааа","аааа","ааааа","аааааа","ааааааа","ааааааааа","аааааааа"}

	res := wordBreak(s, wordDict)
	fmt.Println(res)
}

/*
	139. Word Break
	Medium
	https://leetcode.com/problems/word-break/
	Runtime 0ms Beats 100.00%, Memory 4.17MB Beats 16.89%
 */
func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			test := s[j:i]
			_=test
			// if previous subseq(dp[j]) finished before next
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}