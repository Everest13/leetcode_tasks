package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1,5,0,10,14}

	nums = []int{6,6,0,1,1,4,6} //2

	//nums = []int{1,5,6,14,15} //1

	//nums = []int{82,81,95,75,20} //1


	res := minDifference(nums)
	fmt.Println(res)
}

/*
	1509. Minimum Difference Between Largest and Smallest Value in Three Moves
	Medium
	https://leetcode.com/problems/minimum-difference-between-largest-and-smallest-value-in-three-moves/description/?envType=daily-question&envId=2024-07-03
	Runtime 55ms Beats 93.10%, Memory 7.84MB Beats 48.28%
*/
func minDifference(nums []int) int {
	n:= len(nums)-1
	if n < 4 {
		return 0
	}

	sort.Ints(nums)
	if nums[n] - nums[0] == 0 {
		return 0
	}

	//	min 5 el-v
	minDiff := nums[n-3] - nums[0]
	j:=4
	for j > 0 {
		k:=0
		for k<j {
			diff := nums[n] - nums[k]
			if diff < minDiff {
				minDiff = diff
			}
			k++
		}
		j--
		n--
	}

	return minDiff
}





