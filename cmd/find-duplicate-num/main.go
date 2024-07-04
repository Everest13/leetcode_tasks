package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 3, 4, 2, 2}

	res := findDuplicate(nums)
	fmt.Println(res)
}

/*
	287. Find the Duplicate Number
	Medium
	https://leetcode.com/problems/find-the-duplicate-number/description/?envType=daily-question&envId=2024-05-11
	Runtime 74ms Beats 72.13%, Memory 9.08MB Beats 39.96%
 */
func findDuplicate(nums []int) int {
	n := len(nums)

	mapNums := make(map[int]struct{}, n)
	for i:=0; i<n; i++ {
		if _, ok := mapNums[nums[i]]; ok {
			return nums[i]
		}

		mapNums[nums[i]] = struct{}{}
	}

	return 0
}

// Runtime 93ms Beats 34.22%, Memory 7.08MB Beats 98.96%
func findDuplicate2(nums []int) int {
	n := len(nums)
	sort.Ints(nums)

	for i:=1; i<n; i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}

	return 0
}