package main

import "fmt"

func main() {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	s = "ААААААААААА"
	s = "AAAAAAAAAAAAA"

	res := findRepeatedDnaSequences(s)
	fmt.Println(res)
}

/*
	187. Repeated DNA Sequences
	Medium
	https://leetcode.com/problems/repeated-dna-sequences/description/
	Runtime 32ms, Beats 10.35%, Memory 14.28MB Beats 5.60%
 */
func findRepeatedDnaSequences(s string) []string {
	sArr := []rune(s)
	dnaMap := map[string]int{}
	res := []string{}
	for i:=0; i<len(sArr)-9; i++ {
		dnaMap[string(sArr[i:i+10])]++
	}

	for dna, count := range dnaMap {
		if count > 1 {
			res = append(res, dna)
		}
	}

	return res
}
