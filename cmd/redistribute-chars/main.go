package main

import "fmt"

/*
	1897. Redistribute Characters to Make All Strings Equal
	Easy
	https://leetcode.com/problems/redistribute-characters-to-make-all-strings-equal/description/?envType=daily-question&envId=2023-12-30
	Runtime 11 ms Beats 66.67%, Memory 3.8 MB Beats 88.89%
 */
func main() {
	words := []string{"abc","aabc","bc"}
	words = []string{"abbab"}
	words = []string{"bcc","acaab","bca"}

	res := makeEqual(words)
	fmt.Println("Result is ", res)
}

func makeEqual(words []string) bool {
	lenWords := len(words)
	wordsMap := map[rune]int{}
	for i:=0; i<lenWords; i++ {
		for _, char := range words[i] {
			wordsMap[char]++
		}
	}

	for _, count := range wordsMap {
		if count%lenWords != 0 {
			return false
		}
	}

	return true
}
