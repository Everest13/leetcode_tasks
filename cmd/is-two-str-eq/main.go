package main

import (
	"fmt"
	"sort"
)

/*
	1662. Check If Two String Arrays are Equivalent
	Easy
	https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/description/?envType=daily-question&envId=2023-12-01
	Runtime 4 ms Beats 69.90%, Memory 2.8 MB Beats 92.2%
 */
func main() {
	word1 := []string{"ab", "c"}
	word2 := []string{"a", "bc"}

	word1 = []string{"ab","c"}
	word2 = []string{"a","bcd"}

	res := arrayStringsAreEqual(word1, word2)

	fmt.Println("Res is ", res)
}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	k, m := 0, 0
	for i := 0; i<len(word1); i++ {
		for j := 0; j < len(word1[i]); j++ {
			if k >= len(word2) || word1[i][j] != word2[k][m] {
				return false
			}

			m++
			if m >= len(word2[k]) {
				k++
				m=0
			}
		}
	}

	if k != len(word2) || m != 0 {
		return false
	}

	return true
}


func arrayStringsAreEqualLonger(word1 []string, word2 []string) bool {
	var w1, w2 []byte
	for i := 0; i < len(word1); i++ {
		w1 = append(w1, []byte(word1[i])...)
	}
	for i := 0; i < len(word2); i++ {
		w2 = append(w2, []byte(word2[i])...)
	}

	sort.Slice(w1, func(i, j int) bool { return w1[i] < w1[j] })
	sort.Slice(w2, func(i, j int) bool { return w2[i] < w2[j] })

	return string(w1) == string(w2)
}
