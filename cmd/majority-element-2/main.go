package main

import "fmt"

/*
	229. Majority Element II
	https://leetcode.com/problems/majority-element-ii/description/?envType=daily-question&envId=2023-10-05
	Runtime: 9 ms Beats 71.6%, Memory Usage: 5.1 MB Beats 11.3% M
*/
func main() {
	arr := []int{2,2}

	res := majorityElement(arr)
	fmt.Println("Result is ", res)
}

func majorityElement(nums []int) []int {
	lenNums := len(nums)
	appearMap := map[int]int{}
	for i:=0; i<lenNums; i++ {
		num := nums[i]
		if inc, ok := appearMap[num]; ok {
			appearMap[num] = inc+1
			continue
		}
		appearMap[num] = 1
	}

	times := lenNums/3
	nums = []int{}
	for num, inc := range appearMap {
		if inc > times {
			nums = append(nums, num)
		}
	}

	return nums
}
