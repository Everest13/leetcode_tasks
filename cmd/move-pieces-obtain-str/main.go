package main

import (
	"fmt"
)

func main() {
	banned := []int{1,6,5}
	n := 5
	maxSum := 6

	res := maxCount(banned, n, maxSum)
	fmt.Println(res)
}

/*
	2554. Maximum Number of Integers to Choose From a Range I
	Medium
	https://leetcode.com/problems/maximum-number-of-integers-to-choose-from-a-range-i/description/?envType=daily-question&envId=2024-12-06
	Runtime 27 Beats 100.00%, Memory 9.04MB Beats 80.00%
 */
func maxCount(banned []int, n int, maxSum int) int {
	bannedMap := make(map[int]struct{}, len(banned))
	for _, num := range banned {
		if num <= n {
			bannedMap[num] = struct{}{}
		}
	}

	sum := 0
	count := 0
	for i:=1; i<=n; i++ {
		if _, ok := bannedMap[i]; ok {
			continue
		}
		newSum := sum+i
		if  sum+i > maxSum {
			break
		}
		sum = newSum
		count++
	}

	return count
}