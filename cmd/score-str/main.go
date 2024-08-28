package main

import (
	"fmt"
)

func main() {
	s := "10101"

	res := scoreOfString(s)
	fmt.Println(res)
}

/*
	3110. Score of a String
	Easy
	https://leetcode.com/problems/score-of-a-string/
	Runtime 0ms Beats 100.00%, Memory 2.26MB Beats 91.81%
*/
func scoreOfString(s string) int {
	summ := 0

	for i:=0; i<len(s)-1; i++ {
		diff := int(s[i]) - int(s[i+1])
		if diff < 0 {
			diff *= -1
		}

		summ += diff
	}

	return summ
}