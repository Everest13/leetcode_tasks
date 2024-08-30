package main

import "fmt"

func main() {
	fmt.Println(minPartitions("32"))    // Output: 3
	fmt.Println(minPartitions("82734")) // Output: 8
}

/*
	1689. Partitioning Into Minimum Number Of Deci-Binary Numbers
	Medium
	https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/
	Runtime 6ms Beats 82.03%, Memory 6.42MB Beats67.97%
*/
func minPartitions(n string) int {
	max := 0
	for _, ch := range n {
		digit := int(ch - '0') //convert from int32 to int
		if digit > max {
			max = digit
		}
	}
	return max
}