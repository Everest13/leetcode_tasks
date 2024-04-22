package main

import (
	"fmt"
	"sort"
)

func main() {
	prices := []int{7,6,4,3,1} //0
	prices = []int{7,1,5,3,6,4} //5
	prices = []int{2,1,2,1,0,1,2} //2
	//prices = []int{3,3,5,0,0,3,1,4} //4


	res := maxProfit(prices)
	fmt.Println(res)
}


/*
	121. Best Time to Buy and Sell Stock
	Easy
	https://leetcode.com/problems/best-time-to-buy-and-sell-stock/?envType=study-plan-v2&envId=top-interview-150
	Runtime 92ms Beats 73.08%
	Memory 8.08MB Beats 56.07%
 */
func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	minP := prices[0]
	diff := 0
	for i:=1; i<len(prices); i++ {
		if prices[i] < minP {
			minP = prices[i]
			continue
		}

		if prices[i] > minP {
			locDiff := prices[i] - minP
			if locDiff > diff {
				diff = locDiff
			}
		}
	}

	return diff
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