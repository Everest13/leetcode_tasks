package main

import "fmt"

func main() {
	nums := []int{5,7,7,8,8,10}
	target := 8 //3,4

	nums = []int{5,7,7,8,8,10}
	target = 8 //-

	nums = []int{1}
	target = 1 //1

	//nums = []int{2,2}
	//target = 2 //0,1

	//nums = []int{1,3}
	//target = 1 //0,0
	//
	//nums = []int{1,4}
	//target = 4 //1,1
	//
	//nums = []int{1,2,3}
	//target = 3 //1,1

	res := searchRange(nums, target)
	fmt.Println(res)
}

/*
	34. Find First and Last Position of Element in Sorted Array
	https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/
	Runtime 0ms Beats 100.00%, Memory 6.29MB Beats 7.58%
 */
func searchRange(nums []int, target int) []int {
	n := len(nums)-1
	if n < 0 || target > nums[n] || target < nums[0] {
		return []int{-1,-1}
	}

	start := 0
	mid := (n-start)/2
	for {
		if start == n && nums[start] != target {
			return []int{-1,-1}
		}

		if nums[mid] == target {
			break
		}

		if target > nums[mid] {
			//движемся вправо
			start = mid+1
		} else {
			//движемся влево
			n = mid-1
		}

		mid = n - (n-start)/2
	}



	s := mid
	e := mid
	for {
		count := 0
		if s-1>=0 && nums[s-1]==target {
			s--
			count++
		}

		if e+1<len(nums) && nums[e+1]==target {
			e++
			count++
		}

		if count == 0 {
			break
		}
	}


	return []int{s,e}
}