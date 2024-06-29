package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	low := 58
	high := 155 //[67,78,89,123]

	//low = 1000
	//high = 13000 //[1234,2345,3456,4567,5678,6789,12345]

	//low = 8511
	//high = 23553 //[12345,23456]

	res := sequentialDigits(low, high)
	fmt.Println(res)
}

/*
	1291. Sequential Digits
	Medium
	https://leetcode.com/problems/sequential-digits/description/?envType=daily-question&envId=2024-05-11
	Runtime 1ms Beats 76.19%, Memory 2.19MB Beats 66.67%
 */
func sequentialDigits(low int, high int) []int {
	res := []int{}

	getSeq := func(i, p int) (num int) {
		for p>0 {
			num += i * int(math.Pow10(p-1))
			p--
			i++
		}
		return num
	}

	p := len(strconv.Itoa(low))
	j := low/int(math.Pow10(p-1))
	if getSeq(j,p) < low {
		j++
	}

	num := 0
	for {
		if j>(10-p) {
			j=1
			p++
		}

		num = getSeq(j,p)
		if num > high {
			break
		}

		res = append(res, num)
		j++
	}

	return res
}
