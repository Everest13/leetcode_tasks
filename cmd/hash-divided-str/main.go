package main

import "fmt"

func main() {
	s := "abcd"
	k := 2

	s = "mxz"
	k = 3

	//s = "ao"
	//k = 1

	res := stringHash(s, k)
	fmt.Println(res)
}

/*
	3271. Hash Divided String
	Medium
	https://leetcode.com/problems/hash-divided-string/description/
 */
func stringHash(s string, k int) string {
	n := len(s)
	subLen := n/(n/k)
	res := ""

	for i:=0; i < n; i+=subLen {
		subStr := s[i:i+subLen]
		numStr := 0
		for _, char := range subStr {
			numStr += int(char - 'a')
		}
		hashedChar := rune((numStr % 26) + 'a')
		res += string(hashedChar)
	}

	return res
}