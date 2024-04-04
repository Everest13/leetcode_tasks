package main

import (
	"fmt"
	"sort"
)

func main() {
	//1
	nums1 := []int{1,2,3,0,0,0}
	m := 3
	nums2 := []int{2,5,6}
	n := 3

	//2
	//nums1 = []int{1,0}
	//m = 1
	//nums2 = []int{2}
	//n = 1

	//3
	//nums1 = []int{1}
	//m = 1
	//nums2 = []int{}
	//n = 0

	//4
	//nums1 = []int{2,0}
	//m = 1
	//nums2 = []int{1}
	//n = 1

	merge(nums1, m, nums2, n)
	fmt.Println(nums1)

	nums := []int{3,2,3}
	nums = []int{2,2,1,1,1,2,2}
	nums = []int{3,3,4}
	majorityElement(nums)
	//fmt.Println(nums1)

	nums = []int{1,2,3,4,5,6,7}
	k := 3

	nums = []int{-1}
	k = 2

	nums = []int{1,2}
	k = 3

	rotate(nums, k)
	fmt.Println(nums)
}

/*
	88. Merge Sorted Array
	https://leetcode.com/problems/merge-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150
	Simple
	Runtime 2ms Beats 73.26%, Memory 2.64MB, Beats 14.96%
*/
func merge(nums1 []int, m int, nums2 []int, n int)  {
	if n == 0 {
		nums1 = append(nums1[0:m])
		return
	}

	res := make([]int, 0, m+n)
	j:=0
	for i:=0; i<m; i++ {
		test := []int{}
		for j<n && nums2[j] <= nums1[i] {
			test = append(test, nums2[j])
			j++
		}

		res = append(res, test...)
		res = append(res, nums1[i])
	}

	if j < n {
		res = append(res, nums2[j:]...)
	}

	nums1 = append(nums1[0:0], res...)
}

/*
	169. Majority Element
	https://leetcode.com/problems/majority-element/?envType=study-plan-v2&envId=top-interview-150
	Easy
	Runtime 16ms Beats 58.75%, Memory 6.61MB Beats25.52%
*/
func majorityElement(nums []int) int {
	n := len(nums)
	majoritySign := n / 2

	sort.Ints(nums)
	current := nums[0]
	k := 0
	for i:=1; i<n; i++ {
		if nums[i] != current {
			current = nums[i]
			k=i
			continue
		}

		if i-k+1 > majoritySign {
			return current
		}
	}


	return current
}


/*
	189. Rotate Array
	https://leetcode.com/problems/rotate-array/description/?envType=study-plan-v2&envId=top-interview-150
	Easy
	Runtime 23ms Beats 78.54%, Memory 8.35MB Beats 14.90%
 */
func rotate(nums []int, k int)  {
	l:=len(nums)
	res := make([]int, l)

	if k > l {
		k = k-(l*( k/l))
	}

	res = append(nums[l-k:], nums[0:l-k]...)
	nums = append(nums[0:0], res...)
}