package main

import (
	"fmt"
)

/*
	2125. Number of Laser Beams in a Bank
	Medium
	https://leetcode.com/problems/number-of-laser-beams-in-a-bank/description/?envType=daily-question&envId=2024-01-03
	Runtime 24 ms Beats 84.44%, Memory 7.3 MB Beats 28.89%
 */
func main() {
	test := []string{"011001","000000","010100","001000"}


	res := numberOfBeams(test)
	fmt.Println(res)
}

func numberOfBeams(bank []string) int {
	countSecDevsPrev := 0
	res := 0
	i:=0
	for i<len(bank) {
		countSecDevsCurr := 0
		for _, char := range bank[i] {
			if char == '1' {
				countSecDevsCurr++
			}
		}

		res += countSecDevsPrev * countSecDevsCurr
		i++

		if countSecDevsCurr > 0 {
			countSecDevsPrev = countSecDevsCurr
		}
	}

	return res
}
