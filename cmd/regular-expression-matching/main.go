package main

import (
	"fmt"
	"regexp"
)

/*
	10. Regular Expression Matching
	Hard
	https://leetcode.com/problems/regular-expression-matching/description/
	Runtime - 3ms(beats 64%), memory: 3.8mb(beats 14.9%)
*/
func main() {
	test1 := "aa"
	test2 := "a"

	result := isMatch(test1, test2)
	fmt.Println(result)
}

func isMatch(s string, p string) bool {
	re := regexp.MustCompile(p)
	newSArr := re.FindStringSubmatch(s)

	if len(newSArr) > 0 && newSArr[0] == s {
		return true
	}

	return false
}
