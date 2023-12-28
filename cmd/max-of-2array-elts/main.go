package main

import (
	"fmt"
	"sort"
)

/*
	1464. Maximum Product of Two Elements in an Array
	Easy
	https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/?envType=daily-question&envId=2023-12-12
	Runtime 2 ms Beats 84.16%, Memory 2.6 MB Beats 54.46%
*/
func main() {
	//maxProduct()
	nums := []int{3,4,5,2}
	resnNums := maxProduct(nums)
	fmt.Println(resnNums)
}


func maxProduct(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	res := (nums[0]-1)*(nums[1]-1)

	return res
}
