package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
9. Palindrome Number. Easy level.
https://leetcode.com/problems/palindrome-number/description/
*/
func main() {
	test := 1000021

	result := isPalindrome(test)
	fmt.Println("result is ", result)
}

/*
Runtime - 23ms(beats 63%), memory: 4.8mb(beats 24.29%)
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	sLen := len(strconv.Itoa(x))
	i := 10
	j := int(math.Pow(float64(i), float64(sLen-1)))
	for x > 0 {
		one := x % i
		two := x / j
		if one != two {
			return false
		}

		x = x - two*j
		x = (x - one) / i

		j = j / 100
	}

	return true
}

/*
Runtime - 24ms(beats 55%), memory: 4.8mb(beats 24.29%)
*/
func isPalindrome2(x int) bool {
	s := strconv.Itoa(x)
	lenS := len(s) - 1

	for i := 0; i <= lenS; i++ {
		one := string(s[i])
		two := string(s[lenS-i])
		if one < two && one != two {
			return false
		}
	}

	return true
}
