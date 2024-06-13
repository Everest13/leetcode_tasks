package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "cccccdd" //7
	//s = "a"

	res := longestPalindrome(s)
	fmt.Println(res)
}

/*
	409. Longest Palindrome
	https://leetcode.com/problems/longest-palindrome/description/?envType=daily-question&envId=2024-06-04
	Runtime 2ms Beats 63.35% Memory 2.39MB Beats 63.04% E
 */
func longestPalindrome(s string) int {
	maxPalLen := 0
	charMap := map[rune]int{}
	for _, ch := range s {
		charMap[ch]++
	}

	oneAdded := false
	for _, count := range charMap {
		countDiv := count%2
		if oneAdded == false && countDiv > 0 {
			maxPalLen++
			oneAdded = true
		}

		count = count - (countDiv)
		maxPalLen +=count

	}


	return maxPalLen
}


func longestPalindrome2(s string) int {

	maxPalLen := 0
	sSplited := strings.Split(s, "")
	sort.Strings(sSplited)
	lenS := len(sSplited)

	oneAdded := false
	for i:=0; i<lenS;  {

		count := 1
		for i<lenS-1 && sSplited[i] == sSplited[i+1] {
			count++
			i++
			continue
		}

		countDiv := count%2
		if oneAdded == false && countDiv> 0 {
			maxPalLen++
			oneAdded = true
		}

		count = count - (countDiv)
		maxPalLen += count
		i++
	}


	return maxPalLen
}

