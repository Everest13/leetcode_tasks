package main

import "fmt"

func main() {
	haystack := "a"
	needle := "a"

	res := strStr(haystack, needle)
	fmt.Println(res)
}

/*
	28. Find the Index of the First Occurrence in a String
	Easy
	https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description/
	Runtime 0ms Beats 100.00%, Memory 3.80MB Beats 41.91%
 */
func strStr(haystack string, needle string) int {
	first := needle[0]
	lenNeedle := len(needle)
	lenHaystack := len(haystack)

	for i:=0; i<=lenHaystack-lenNeedle; i++ {
		if haystack[i] == first && haystack[i:i+lenNeedle] == needle {
			return i
		}
	}

	return -1
}
