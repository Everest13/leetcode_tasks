package main

import (
	"fmt"
)

/*
	1930. Unique Length-3 Palindromic Subsequences
	https://leetcode.com/problems/unique-length-3-palindromic-subsequences/description/?envType=daily-question&envId=2023-11-14
	Runtime 288 ms Beats 88.89%, Memory 6.7 MB Beats 66.67% M
 */
func main() {
	test := "bbcbaba"
	//test = "tsltpjzdmtwderpkpmgoyrcxttiheassztncqvnfjeyxxp"
	//test = "www"
	test = "aabca"
	res := countPalindromicSubsequence(test)
	fmt.Println("Result is ", res)
}

type occurrence struct {
	first int
	last int
}

//Faster in 3 times
func countPalindromicSubsequence(s string) int {
	var res int
	chIDxMap := map[rune]occurrence{}
	for i, ch := range s {
		if occ, ok := chIDxMap[ch]; ok {
			occ.last = i
			chIDxMap[ch] = occ
			continue
		}

		chIDxMap[ch] = occurrence{first: i}
	}

	r := []rune(s)
	for _, occ := range chIDxMap {
		if occ.last == 0 {
			continue
		}
		uniqCh := map[rune]struct{}{}
		for i:=occ.first+1; i<occ.last; i++ {
			uniqCh[r[i]] = struct{}{}
		}

		res += len(uniqCh)
	}

	return res
}





// First Decision
func countPalindromicSubsequenceFirst(s string) int {
	polindromicMap := map[string]struct{}{}
	chIDxMap := map[rune]int{}
	tripleCh := map[rune]int{}

	r := []rune(s)
	for i, ch := range s {
		tripleCh[ch]++

		if idx, ok := chIDxMap[ch]; ok {
			if tripleCh[ch] > 2 {
				polindromicMap[string([]rune{ch, ch, ch})] = struct{}{}
			}

			if i==idx+1{
				continue
			}

			for j:=idx+1; j<i; j++ {
				polindromicMap[string([]rune{ch, r[j], ch})] = struct{}{}
			}

			chIDxMap[ch] = i
			continue
		}

		chIDxMap[ch] = i
	}

	return len(polindromicMap)
}

