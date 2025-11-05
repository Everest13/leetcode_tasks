package main

func main() {
	nums := []int{1,1,2}
	nums = []int{0,0,1,1,1,2,2,3,3,4}

	res := removeDuplicates(nums)
	_=res
}

/*
	26. Remove Duplicates from Sorted Array
	Easy
	https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
 */
func removeDuplicates(nums []int) int {
	j:=0
	for i:=1; i<len(nums); i++ {
		if nums[j] == nums[i] {
			continue
		}
		j++
		nums[j] = nums[i]
	}

	j++
	return j
}
