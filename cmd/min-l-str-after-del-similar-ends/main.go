package main

import "fmt"

func main() {

	s := "aabccabba" // 3
	//s = "ca" //2

	//s = "cabaabac" //0
	//s = "abbbbbbbbbbbbbbbbbbba" // 0

	//s = "bbbbbbbbbbbbbbbbbbbbbbbbbbbabbbbbbbbbbbbbbbccbcbcbccbbabbb" //1
	// abbbbbbbbbbbbbbbccbcbcbccbba
	// bbbbbbbbbbbbbbbccbcbcbccbb
	// ccbcbcbcc
	// bcbcb
	// cbc
	// b

	//s = "cbc"

	res := minimumLength(s)

	fmt.Println(res)
}

/*
	1750. Minimum Length of String After Deleting Similar Ends
	https://leetcode.com/problems/minimum-length-of-string-after-deleting-similar-ends/description/?envType=daily-question&envId=2024-03-05
	Runtime 5ms Beats 78.68%
 */
func minimumLength(s string) int {
	lenS :=len(s)
	if  lenS < 2 {
		return lenS
	}

	i:=0
	j := lenS - 1
	for s[i] == s[j] {
		if j-i < 3 {
			return j-i-1
		}

		i++
		j--

		for i+1 < j && s[i] == s[i-1] {
			i++
		}
		for i+1 < j && s[j] == s[j+1] {
			j--
		}
	}

	return j-i+1
}

