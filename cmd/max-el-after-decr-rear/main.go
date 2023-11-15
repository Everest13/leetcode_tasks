package main

import (
	"fmt"
	"sort"
)

/*
	1846. Maximum Element After Decreasing and Rearranging
	Medium
	https://leetcode.com/problems/maximum-element-after-decreasing-and-rearranging/description/?envType=daily-question&envId=2023-11-15
	Runtime 76 ms Beats 75%, Memory 8.7 MB Beats 50%
 */
func main() {
	arr := []int{2,2,1,2,1}
	arr=[]int{91205}

	res := maximumElementAfterDecrementingAndRearranging(arr)
	fmt.Println("Res is ", res)
}

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)

	max := 1
	checkMax := func(i int) {
		if i > max{
			max = i
		}
	}

	arr[0] = 1
	lenArr := len(arr)
	for i:=1; i < lenArr; i++ {
		diff := arr[i-1] - arr[i]
		if diff <= 1 && diff >= -1 {
			checkMax(arr[i])
			continue
		}

		arr[i] = arr[i-1] + 1
		checkMax(arr[i])
	}

	return max
}
