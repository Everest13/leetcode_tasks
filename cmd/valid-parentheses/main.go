package main

import "fmt"

var bracketMap = map[string]string{
	")": "(",
	"}": "{",
	"]": "[",
}

/*
Task #6 Valid Parentheses. Easy level.
Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
Memory Usage: 2.1 MB, less than 24.85% of Go online submissions for Valid Parentheses.
https://leetcode.com/problems/valid-parentheses/
*/
func main() {
	brackets := "()" //"{[()]}{}()[][{()}]"
	result := isValid(brackets)

	fmt.Println("Result is ", result)
}

func isValid(s string) bool {
	pocket := make([]string, 0, len(s))
	var str string
	for i := 0; i < len(s); i++ {
		str = string(s[i])
		if str == "{" || str == "(" || str == "[" {
			pocket = append(pocket, str)
			continue
		}

		if len(pocket) == 0 {
			return false
		}
		if openBracket, ok := bracketMap[str]; ok && openBracket == pocket[len(pocket)-1] {
			pocket = pocket[:len(pocket)-1]
			continue
		}

		return false
	}

	if len(pocket) > 0 {
		return false
	}

	return true
}
