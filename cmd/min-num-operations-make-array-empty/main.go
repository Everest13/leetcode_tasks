package main

import (
	"fmt"
	"sort"
)

/*
	2870. Minimum Number of Operations to Make Array Empty
	Medium
	https://leetcode.com/problems/minimum-number-of-operations-to-make-array-empty/description/?envType=daily-question&envId=2024-01-04
	Runtime 105 ms Beats 85.19%, Memory 8.8 MB Beats 92.59%
 */
func main() {
	test := []int{2,3,3,2,2,4,2,3,4}
	test = []int{14,12,14,14,12,14,14,12,12,12,12,14,14,12,14,14,14,12,12}

	res := minOperations(test)
	fmt.Println(res)
}

func minOperations(nums []int) int {
	sort.Ints(nums)

	opNums := 0
	addOpt := func(count int) {
		// 0- кратно трем, 1- один лишний - можно заменить на 2е двойки (+1), 2- одна двойка (+1)
		switch count%3 {
		case 0:
			opNums += count / 3
		case 1, 2:
			opNums += count/3 + 1
		}
	}

	count := 1
	for i:=0; i<len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			count++
			continue
		}

		if count < 2 {
			return -1
		}

		addOpt(count)
		count=1
	}

	if count < 2 {
		return -1
	}

	addOpt(count)

	return opNums
}
