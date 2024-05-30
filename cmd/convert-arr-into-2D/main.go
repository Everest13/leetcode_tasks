package main

import (
	"fmt"
	"sort"
)

/*
	2610. Convert an Array Into a 2D Array With Conditions
	https://leetcode.com/problems/convert-an-array-into-a-2d-array-with-conditions/description/?envType=daily-question&envId=2024-01-02
	Runtime 3 ms Beats 76%, Memory 3.3 MB Beats 96% M
 */
func main() {
	test := []int{1,3,4,1,2,3,1}
	test = []int{9,8,8,4,9,8,8,3,9}


	res := findMatrix(test)
	fmt.Println(res)
}

func findMatrix(nums []int) [][]int {
	sort.Ints(nums)

	count := 1
	res := [][]int{0: {nums[0]}}
	for i:=1; i<len(nums); i++ {
		if nums[i-1] == nums[i] {
			if len(res)-1 < count {
				res = append(res, []int{nums[i]})
				count++
				continue
			}

			res[count] = append(res[count], nums[i])
			count++
			continue
		}

		res[0] = append(res[0],  nums[i])
		count=1
	}

	return res
}

func findMatrix2(nums []int) [][]int {
	res := [][]int{0: {}}
	distinctMap := map[int]int{}

	for _, num := range nums {
		if count, ok := distinctMap[num]; ok {
			if len(res)-1 < count {
				res = append(res, []int{num})
				distinctMap[num]++
				continue
			}
			res[count] = append(res[count], num)
			distinctMap[num]++
			continue
		}

		res[0] = append(res[0], num)
		distinctMap[num]++
	}

	return res
}
