package main

import (
	"fmt"
	"math"
)

/*
7. Reverse Integer. Medium level.
https://leetcode.com/problems/reverse-integer/submissions/915535680/
Runtime - 0ms(beats 100%), memory: 2.1mb(beats 16.9%)
*/
func main() {
	test := -123

	result := reverse(test)
	fmt.Println(result)
}

func reverse(x int) int {
	result := 0
	for x != 0 {
		result = result*10 + x%10
		x = x / 10
	}

	if x < 0 {
		result = 0 - result
	}

	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}

	return result
}
