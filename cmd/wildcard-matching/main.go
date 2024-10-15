package main

import (
	"fmt"
)

func main() {
	s := "acdcfb"

	p := "a*c?b" //true

	//s = "cb"
	//p = "?a" //false

	//s = ""
	//p = "***"

	s = "a"
	p = "aa" //false

	s = "zacabz"
	p = "*a?b*" //false

	s = "aaaa"
	p = "***a"

	//s = "a"
	//p = "?*"

	res := isMatch(s, p)
	fmt.Println(res)
}

/*
	44. Wildcard Matching
	Hard
	https://leetcode.com/problems/wildcard-matching/description/
	Runtime 16ms Beats 42.96%, Memory 6.55MB Beats 62.22%
 */

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)

	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-1]
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			t1:= string(p[j-1])
			t2 := string(s[i-1])
			_,_ = t1,t2

			if p[j-1] == '?' || p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}

	return dp[m][n]
}
