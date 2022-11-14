package main

import (
	"errors"
	"fmt"
	"strings"
)

// Task #14 Longest Common Prefix
// https://leetcode.com/problems/longest-common-prefix/submissions/
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Common Prefix.
// Memory Usage: 2.2 MB, less than 86.62% of Go online submissions for Longest Common Prefix.
func main() {
	strs := []string{"flower", "flow", "flight"}
	err := validateStrs(strs)
	if err != nil {
		fmt.Println(err.Error())
	}

	result := getLongestCommonPrefix(strs)
	if len(result) == 0 {
		fmt.Println(errors.New("there is no common prefix among the input strings"))
		return
	}

	fmt.Println("Result is ", result)
}

func validateStrs(strs []string) error {
	if len(strs) > 200 || len(strs) == 0 {
		return errors.New("array of strings contain invalid number of elements")
	}

	for i, str := range strs {
		strs[i] = strings.ToLower(str)
		if len(str) > 200 || len(str) == 0 {
			return errors.New("string contain invalid number of characters")
		}
	}

	return nil
}

// O(n^2)
func getLongestCommonPrefix(strs []string) string {
	var prefix string
	checkingStr := strs[0]
	strs = strs[1:]
	for _, char := range checkingStr {
		prefix += string(char)
		for _, str := range strs {
			if !strings.HasPrefix(str, prefix) {
				return prefix[:len(prefix)-1]
			}
		}
	}

	return prefix
}
