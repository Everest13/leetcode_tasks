package main

import "fmt"

func main() {
	t := "azw"

	res := stringSequence(t)
	fmt.Println(res)
}

/*
	3324. Find the Sequence of Strings Appeared on the Screen
	Medium
	https://leetcode.com/problems/find-the-sequence-of-strings-appeared-on-the-screen/
	Runtime 6ms Beats 98.80%, Memory 19.48MB Beats 49.40%
 */
func stringSequence(target string) []string {
	n := len(target)
	totalSize := 0

	for i := 0; i < n; i++ {
		totalSize += int(target[i] - 'a' + 1)
	}

	res := make([]string, totalSize)
	subseq := make([]byte, 0, n)

	index := 0

	for i := 0; i < n; i++ {
		tI := target[i]
		start := len(subseq)

		for l := uint8('a'); l <= tI; l++ {
			subseq = append(subseq[:start], l)
			res[index] = string(subseq)
			index++
		}
	}

	return res
}
