package main

import "fmt"

func main() {
	s := "aaaaaaabccbaa" //4
	k := 2

	//s = "abc" //3
	//k = 1

	//s = "acba" //3
	//k = 1

	res := takeCharacters(s, k)
	fmt.Println(res)
}

/*
	2516. Take K of Each Character From Left and Right
	Medium
	https://leetcode.com/problems/take-k-of-each-character-from-left-and-right/description/
	Runtime 0ms Beats 100.00%, Memory 7.63MB Beats 36.18%
 */
func takeCharacters(s string, k int) int {
	if k == 0 {
		return 0
	}

	n := len(s)
	total := [3]int{} // Indexes: 0 -> 'a', 1 -> 'b', 2 -> 'c'
	for i := 0; i < n; i++ {
		total[s[i]-'a']++
	}

	if total[0] < k || total[1] < k || total[2] < k {
		return -1
	}

	// How many characters of each type can remain inside the window
	// so that the remaining characters outside the window satisfy the problem condition
	target := [3]int{
		total[0] - k,
		total[1] - k,
		total[2] - k,
	}

	// Sliding window: find the maximum window containing target symbols
	left, maxWindow := 0, 0
	current := [3]int{}

	for right := 0; right < n; right++ {
		current[s[right]-'a']++ // increment the current symbol

		// Shrink the window if it has too many characters
		for current[0] > target[0] || current[1] > target[1] || current[2] > target[2] {
			current[s[left]-'a']--
			left++
		}

		// Remember the maximum window size
		maxWindow = max(maxWindow, right-left+1)
	}

	return n - maxWindow
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
