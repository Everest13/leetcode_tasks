package main

import (
	"fmt"
	"strings"
)

type node struct {
	val int
	s   string
}

var romanNodes = []node{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// Other solution
// https://leetcode.com/problems/integer-to-roman/discuss/625791/Golang-8ms
func main2() {
	// input integer <= 3999
	inputInt := 2348
	result := intToRoman2(inputInt)

	fmt.Println("Result is ", result)
}

func intToRoman2(num int) string {
	sb := strings.Builder{}

	for _, n := range romanNodes {
		for num >= n.val {
			sb.WriteString(n.s)
			num -= n.val
		}
	}

	return sb.String()
}
