package main

import (
	"fmt"
)

/*
	1347. Minimum Number of Steps to Make Two Strings Anagram
	Medium
	https://leetcode.com/problems/minimum-number-of-steps-to-make-two-strings-anagram/description/?envType=daily-question&envId=2024-01-13
	Runtime 28ms Beats 51.16%, Memory 8.49MB Beats 6.98%
 */
func main() {
	s := "leetcode"
	t := "practice"

	//s = "bab"
	//t = "aba"

	res := minSteps(s, t)
	fmt.Println(res)
}

func minSteps(s string, t string) int {
	mapS := make(map[rune]int, len(s))
	for _, char := range s {
		mapS[char]++
	}

	count := 0
	for _, char := range t {
		if amount, ok := mapS[char]; !ok || amount == 0 {
			count++
			continue
		}

		mapS[char]--
	}

	return count
}
