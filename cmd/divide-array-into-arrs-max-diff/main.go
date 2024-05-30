package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1,3,4,8,7,9,3,5,1}
	k := 2

	res := divideArray(nums, k)
	fmt.Println(res)
}

/*
	2966. Divide Array Into Arrays With Max Difference
	https://leetcode.com/problems/divide-array-into-arrays-with-max-difference/description/?envType=daily-question&envId=2024-02-01
	Runtime 155ms beats 80%, Memory 10.6 beats 48% M
 */
func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)

	if len(nums)%3 != 0 {
		return [][]int{}
	}

	res := [][]int{}
	i:=0
	for i<len(nums)-2 {
		first := nums[i]
		third := nums[i+2]
		if third - first > k {
			return [][]int{}
		}

		res = append(res, []int{first, nums[i+1], third})
		i+=3
	}

	return res
}
