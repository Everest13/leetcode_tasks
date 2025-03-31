package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 886996

	res := alternateDigitSum(n)
	fmt.Println(res)
}

/*
	2544. Alternating Digit Sum
	Easy
	https://leetcode.com/problems/alternating-digit-sum/description/
	Runtime 0ms Beats 100.00%, Memory 3.90MB Beats 18.18%
 */
func alternateDigitSum(n int) int {
	str := strconv.Itoa(n)
	res := 0
	sign := 1

	for _, char := range str {
		res += sign*int(char - '0')
		sign *= -1
	}

	return res
}
