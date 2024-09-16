package main

import "fmt"

func main() {
	s := ")()())" //4
	s = "(()" //2
	s = "()(())" //6

	//s = "((()())(())"

	res := longestValidParentheses(s)
	fmt.Println(res)
}

/*
	32. Longest Valid Parentheses
	Hard
	https://leetcode.com/problems/longest-valid-parentheses/
	Runtime 0ms Beats 100.00%, Memory 4.80MB Beats 9.02%
 */

func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}

	stack := []int{}
	idxs := [][2]int{}
	for i:=0; i<len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
			continue
		}
		n := len(stack)
		if n == 0 {
			continue
		}

		idxs = append(idxs, [2]int{i, stack[n-1]})
		stack = stack[:n-1]
	}

	lenIdx := len(idxs)
	if lenIdx == 0 {
		return 0
	}

	lenIdx--
	count := idxs[lenIdx][0] - idxs[lenIdx][1] + 1
	maxCount := count
	lastIdx := idxs[lenIdx][1]
	for i:=lenIdx-1; i>=0; i-- {
		start := idxs[i][1]
		end := idxs[i][0]

		if start > lastIdx {
			continue
		}

		if lastIdx - end == 1 {
			count += end - start + 1
		} else {
			count = end - start + 1
		}

		lastIdx = start
		if maxCount < count {
			maxCount = count
		}
	}

	return maxCount
}