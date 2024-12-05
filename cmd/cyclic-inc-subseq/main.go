package main

import "fmt"

func main() {
	str1 := "abc"
	str2 := "ad" //true

	str1 = "zc"
	str2 = "ad" //true

	//str1 = "ab"
	//str2 = "d" //false

	res := canMakeSubsequence(str1, str2)
	fmt.Println(res)
}

/*
	2825. Make String a Subsequence Using Cyclic Increments
	Medium
	https://leetcode.com/problems/make-string-a-subsequence-using-cyclic-increments/
	Runtime 0ms Beats 100.00%, Memory 8.35MB Beats 11.11%
 */
func canMakeSubsequence(str1 string, str2 string) bool {
	len1 := len(str1)
	len2 := len(str2)
	j:=0
	for i:=0; i < len1; i++ {
		if j==len2 {
			return true
		}
		if str2[j]-str1[i] <= 1 || (str1[i] == 122 && str2[j] == 97) {
			j++

		}
	}

	return j==len2
}
