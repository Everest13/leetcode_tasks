package main

import "fmt"

func main() {
	test := "001011"

	res := minOperations(test)
	fmt.Println(res)
}

/*
	1769. Minimum Number of Operations to Move All Balls to Each Box
	Medium
	https://leetcode.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box/description/
 */
/*
	O(n^2)
 */
func minOperations(boxes string) []int {
	answers := make([]int, len(boxes))

	for i:=0; i<len(boxes); i++ {
		if boxes[i] == 49 {
			for j:=0; j<len(answers); j++ {
				num := i-j
				if num < 0 {
					num*=-1
				}
				answers[j] += num
			}
		}
	}

	return answers
}

/*
	O(n)
*/
func minOperations2(boxes string) []int {
	n := len(boxes)
	res := make([]int, n)

	steps, count := 0, 0
	for i := 0; i < n; i++ {
		res[i] += steps
		if boxes[i] == '1' {
			count++
		}
		steps += count
	}

	steps, count = 0, 0
	for i := n - 1; i >= 0; i-- {
		res[i] += steps
		if boxes[i] == '1' {
			count++
		}
		steps += count
	}

	return res
}