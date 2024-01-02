package main

import "fmt"

func main() {
	s := "abca"
	s = "cbzxy"
	s = "mgntdygtxrvxjnwksqhxuxtrv"
	dst := maxLengthBetweenEqualCharacters(s)
	fmt.Println(dst)
}

/*
	1624. Largest Substring Between Two Equal Characters
	Easy
	https://leetcode.com/problems/largest-substring-between-two-equal-characters/description/?envType=daily-question&envId=2023-12-31
	Runtime 1 ms Beats 70.59%, Memory 2.1 MB Beats 47.6%
*/
func maxLengthBetweenEqualCharacters(s string) int {
	subStrMap := map[rune]int{}
	res := -1
	for i, char := range s {
		if j, ok := subStrMap[char]; ok {
			betw := i - j - 1
			if betw > res {
				res = betw
			}
			continue
		}

		subStrMap[char] = i
	}

	return res
}