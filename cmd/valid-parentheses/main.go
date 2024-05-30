package main

import "fmt"

var bracketMap = map[string]string{
	")": "(",
	"}": "{",
	"]": "[",
}

/*
	20. Valid Parentheses
	https://leetcode.com/problems/valid-parentheses/
	Runtime: 0 ms, faster than 100.00% Memory Usage: 2.1 MB, less than 24.85% E
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
