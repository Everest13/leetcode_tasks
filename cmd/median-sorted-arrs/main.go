package main

import (
	"fmt"
	"math"
)

/*
Task #13 Median of Two Sorted Arrays (Hard)
https://leetcode.com/problems/median-of-two-sorted-arrays/description/
Runtime: 17 ms, faster than 55.00% of Go online submissions for Merge Two Sorted Lists.
Memory Usage: 6.5 MB, less than 12.59% of Go online submissions for Merge Two Sorted Lists.

Offtop:
Not  O(log (m+n)) decision.
*/
func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	result := findMedianSortedArrays(nums1, nums2)
	fmt.Println(result)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var resultArr []int
	len1 := len(nums1)
	len2 := len(nums2)
	i, j := 0, 0

	countMedian := func(resultArr []int) float64 {
		lenArr := len(resultArr)
		lenHalf := lenArr / 2
		if lenArr%2 > 0 { //если не четное число
			return math.Ceil(float64(resultArr[lenHalf]))
		}

		test := (float64(resultArr[lenHalf] + resultArr[lenHalf-1])) / 2
		return test
	}

	for true {
		if i >= len1 {
			resultArr = append(resultArr, nums2[j:]...)
			return countMedian(resultArr)
		}
		if j >= len2 {
			resultArr = append(resultArr, nums1[i:]...)
			return countMedian(resultArr)
		}

		if nums1[i] < nums2[j] {
			resultArr = append(resultArr, nums1[i])
			i++
		} else {
			resultArr = append(resultArr, nums2[j])
			j++
		}
	}

	return 0
}
