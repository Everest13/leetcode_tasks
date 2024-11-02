package main

import "fmt"

func main() {
	nums := []int{1,1,2,3,3,3,4,4,5}
	//nums = []int{1,1,2,2,2,2,2,5,6,7,7,8,8,8,9}
	//nums = []int{0,0,1,1,1,2,2,3,3,4}

	k := removeDuplicates(nums)
	expectedNums := make([]int, 0, k)
	for i:=0; i<=k; i++ {
		expectedNums = append(expectedNums, nums[i])
	}

	fmt.Println(expectedNums)
}

/*
	26. Remove Duplicates from Sorted Array
	Easy
	https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
	Runtime 0ms Beats 100.00%, Memory 6.24MB Beats 5.83%
 */
func removeDuplicates(nums []int) int {
	n := len(nums)
	i:=1
	k:=1
	uniq := nums[i-1]
	for i<n {
		if nums[i] != uniq {
			nums[k] = nums[i]
			uniq = nums[i]
			k++
			i++
			continue
		}
		i++
	}

	return k
}
