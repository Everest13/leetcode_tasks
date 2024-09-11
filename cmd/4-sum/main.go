package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1,0,-1,0,-2,2}
	target := 0

	//nums = []int{1,2,3,4,5}
	//target = 10

	res := fourSum(nums, target)
	fmt.Println(res)
}

/*
	18. 4Sum
	Medium
	https://leetcode.com/problems/4sum/
	Runtime 120ms Beats 15.12%, Memory 7.83MB Beats 8.78%
 */
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	numsMap := map[int]int{}
	for i:=0; i<n; i++ {
		numsMap[nums[i]] = i
	}

	res := [][]int{}
	checkQdrt := func(i, j, k, m int) {
		arr := []int{nums[i], nums[j], nums[k], nums[m]}
		for _, resQdrt := range res {
			if resQdrt[0] == arr[0] && resQdrt[1] == arr[1] && resQdrt[2] == arr[2] && resQdrt[3] == arr[3] {
				return
			}
		}

		res = append(res, arr[:])
	}

	for i:=0; i<n-3; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				lastNum := target - nums[i] - nums[j] - nums[k]
				if m, ok := numsMap[lastNum]; ok && m > k {
					checkQdrt(i, j, k, m)
				}
			}
		}
	}

	return res
}

/*
	Runtime 1418ms Beats 5.22%, Memory 3.06MB Beats 29.35%
 */
func fourSum2(nums []int, target int) [][]int {
	n := len(nums)
	resMap := map[[4]int]struct{}{}
	res := [][]int{}

	sort.Ints(nums)

	checkQuadrplt := func(i, j, k, m int) {
		summ := nums[i] + nums[j] + nums[k] + nums[m]
		if summ != target {
			return
		}

		arr := [4]int{nums[i], nums[j], nums[k], nums[m]}
		if _, ok := resMap[arr]; ok {
			return
		}
		resMap[arr] = struct{}{}
		qdpt := []int{nums[i], nums[j], nums[k], nums[m]}
		res = append(res, qdpt)
	}

	for i:=0; i<n-3; i++ {
		for j:=i+1; j<n; j++ {
			for k:=j+1; k<n; k++ {
				for m:=k+1; m<n; m++ {
					checkQuadrplt(i, j, k, m)
				}
			}
		}
	}

	return res
}

