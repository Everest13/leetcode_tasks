package main

import (
	"fmt"
	"sort"
)

func main() {
	str := []string{"eat","tea","tan","ate","nat","bat"}

	res := groupAnagrams(str)
	fmt.Println(res)

}

/*
	49. Group Anagrams
	Medium
	https://leetcode.com/problems/group-anagrams/description/?envType=daily-question&envId=2024-02-06
	Runtime 14ms Beats 92.27%
	Memory 7.57MB Beats 69.64%
 */
func groupAnagrams(strs []string) [][]string {
	if len(strs) < 2 {
		return [][]string{strs}
	}

	res := [][]string{}
	anagramMap := map[string]int{}
	k:=0
	for _, str := range strs {
		chars := []rune(str)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})

		charsStr := string(chars)
		if idx, ok := anagramMap[charsStr]; ok {
			res[idx] = append(res[idx], str)
			continue
		}

		res = append(res, []string{str})
		anagramMap[charsStr] = k
		k++
	}

	return res
}
