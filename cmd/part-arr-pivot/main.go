package main

import "fmt"

func main() {
	nums := []int{9,12,5,10,14,3,10}
	pivot := 10

	res := pivotArray(nums, pivot)
	fmt.Println(res)
}

/*
	2161. Partition Array According to Given Pivot
	Medium
	https://leetcode.com/problems/partition-array-according-to-given-pivot/description/
 */
func pivotArray(nums []int, pivot int) []int {
	res := make([]int, len(nums))
	i := 0
	j := len(nums)-1
	countPivots := 0
	for _, num := range nums {
		if num < pivot {
			res[i] = num
			i++
			continue
		}
		if num == pivot {
			countPivots++
			continue
		}
		res[j] = num
		j--
	}

	for countPivots > 0 {
		res[i] = pivot
		countPivots--
		i++
	}

	j = len(nums)-1
	for i < j {
		resI := res[i]
		res[i] = res[j]
		res[j] = resI
		i++
		j--
	}

	return res
}
