package main

import "fmt"

func main() {
	nums := []int{4,3,2,7,8,2,3,1}

	res := findDuplicates(nums)
	fmt.Println(res)
}

/*
	442. Find All Duplicates in an Array
	https://leetcode.com/problems/find-all-duplicates-in-an-array/description/?envType=daily-question&envId=2024-03-25
	Medium
	Runtime 41ms Beats 76.56%, Memory 7.59MB Beats 38.28%
 */
func findDuplicates(nums []int) []int {
	res := []int{}
	numsMap := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		if _, ok := numsMap[num]; ok {
			res = append(res, num)
			continue
		}
		numsMap[num] = struct{}{}
	}

	return res
}