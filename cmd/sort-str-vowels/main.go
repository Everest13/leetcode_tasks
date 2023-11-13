package main

import (
	"fmt"
	"sort"
)

/*
	2785. Sort Vowels in a String
	Medium
	https://leetcode.com/problems/sort-vowels-in-a-string/description/?source=submission-noac
	Runtime 25 ms Beats 75.68%; Memory 10.3 MB Beats 27.3%
 */
func main() {
	test := "RiQYo"
	res := sortVowels(test)

	fmt.Println("Result is ", res)
}

var vovels = map[rune]struct{}{
	'a': {},
	'e': {},
	'i': {},
	'o': {},
	'u': {},
	'A': {},
	'E': {},
	'I': {},
	'O': {},
	'U': {},
}

//Все согласные остаются на своих прежних местах.
//Гласные должны быть отсортированы в порядке неубывания их значений ASCII.
func sortVowels(s string) string {
	var vovelS []rune
	for _, ch := range s {
		if _, ok := vovels[ch]; ok {
			vovelS = append(vovelS, ch)
		}
	}

	sort.Slice(vovelS, func(i, j int) bool {
		return vovelS[i] < vovelS[j]
	})

	r := []rune(s)
	j := 0
	for i, ch := range r {
		if _, ok := vovels[ch]; ok {
			r[i] = vovelS[j]
			j++
		}
	}


	return string(r)
}
