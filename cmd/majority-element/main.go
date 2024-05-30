package main

import "sort"

func main() {

}

/*
	169. Majority Element
	https://leetcode.com/problems/majority-element/?envType=study-plan-v2&envId=top-interview-150
	Runtime 16ms Beats 58.75%, Memory 6.61MB Beats25.52% E
*/
func majorityElement(nums []int) int {
	n := len(nums)
	majoritySign := n / 2

	sort.Ints(nums)
	current := nums[0]
	k := 0
	for i:=1; i<n; i++ {
		if nums[i] != current {
			current = nums[i]
			k=i
			continue
		}

		if i-k+1 > majoritySign {
			return current
		}
	}


	return current
}
