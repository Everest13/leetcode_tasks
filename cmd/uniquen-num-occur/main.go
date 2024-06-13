package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1,2}
	arr = []int{1,2,2,1,1,3}

	res := uniqueOccurrences(arr)
	fmt.Println(res)
}

/*
	1207. Unique Number of Occurrences
	https://leetcode.com/problems/unique-number-of-occurrences/?envType=daily-question&envId=2024-05-11
	Runtime 0ms Beats 100.00%, Memory 2.46MB Beats 99.33% E
 */
func uniqueOccurrences(arr []int) bool {
	sort.Ints(arr)
	occurMap := map[int]int{}

	i:=0
	count := 1
	for i<len(arr) {
		count = 1
		for i+1<len(arr) && arr[i] == arr[i+1] {
			count++
			i++
			continue
		}

		if _, ok := occurMap[count]; ok {
			return false
		}

		occurMap[count] = arr[i]
		i++
	}

	return true
}
