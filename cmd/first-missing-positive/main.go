package main

import "fmt"

func main() {
	nums := []int{1,2,0} //3
	nums = []int{3,4,-1,1} //2
	nums = []int{7,8,9,11,12}//1
	nums = []int{3,2}

	res := firstMissingPositive(nums)
	fmt.Println(res)
}

/*
	41. First Missing Positive
	Hard
	https://leetcode.com/problems/first-missing-positive/
	Runtime 44ms Beats 94.22%, Memory 8.52MB Beats 50.87%
 */
func firstMissingPositive(nums []int) int {
	lenN := len(nums)
	arr := make([]int, lenN)
	for i:=0; i<lenN; i++ {
		if nums[i]>0 && nums[i]<=len(arr) {
			arr[nums[i]-1] = 1
			continue
		}
		arr = arr[:len(arr)-1]
	}

	lenArr := len(arr)
	for i:=0; i<lenArr; i++ {
		if arr[i] == 0 {
			return i+1
		}
	}

	return lenArr+1
}
