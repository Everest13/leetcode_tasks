package main

import "fmt"

func main() {
	pref := []int{5,2,0,3,1}

	res := findArray(pref)
	fmt.Println(res)
}

/*
	2433. Find The Original Array of Prefix Xor
	Medium
	https://leetcode.com/problems/find-the-original-array-of-prefix-xor/description/
 */
func findArray(pref []int) []int {
	res := make([]int, len(pref))
	res[0] = pref[0]

	for i:=1; i<len(pref); i++ {
		res[i] = pref[i-1] ^ pref[i]
	}

	return res
}