package main

import (
	"fmt"
)

/*
Task #1 Two Sum.
https://leetcode.com/problems/two-sum/description/
Runtime - 27ms(beats 29%), memory: 4.5mb(beats 85%)
*/
func main() {
	nums := []int{-10, -1, -18, -19}
	target := -19
	result := twoSum(nums, target)
	if len(result) < 2 {
		fmt.Println("Indexes not found")
		return
	}

	fmt.Println("Result is ", result)
}

func twoSum(nums []int, target int) []int {
	lenNums := len(nums)

	for i := 0; i < lenNums; i++ {
		y := target - nums[i]
		j := lenNums - 1

		for j > i {
			if nums[j] == y {
				return []int{i, j}
			}
			j--
		}
	}

	return nil
}

/*
Runtime - 11ms(beats 65%), memory: 4.5mb(beats 85%)
Дает всего одну проходку
*/
func bestSolution(nums []int, target int) []int {
	addends := make(map[int]int)
	for i, num := range nums {
		if j, ok := addends[target-num]; ok {
			return []int{i, j}
		}
		addends[num] = i
	}

	return nil
}
