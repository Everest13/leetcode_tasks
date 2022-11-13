package main

import (
	"errors"
	"fmt"
)

var intRoman = map[int]string{
	3000: "MMM",
	2000: "MM",

	1000: "M",
	900:  "CM",
	800:  "DCCC",
	700:  "DCC",
	600:  "DC",
	500:  "D",
	400:  "CD",
	300:  "CCC",
	200:  "CC",
	100:  "C",

	90: "XC",
	80: "LXXX",
	70: "LXX",
	60: "LX",
	50: "L",
	40: "XL",
	30: "XXX",
	20: "XX",
	10: "X",

	9: "IX",
	8: "VIII",
	7: "VII",
	6: "VI",
	5: "V",
	4: "IV",
	3: "III",
	2: "II",
	1: "I",
}

// Task #2
// Given an integer numeral, convert it to roman.
// https://leetcode.com/problems/integer-to-roman/
// Runtime: 11 ms, faster than 80.96% of Go online submissions for Integer to Roman.
// Memory Usage: 3.1 MB, less than 92.73% of Go online submissions for Integer to Roman.
func main() {
	// input integer <= 3999
	inputInt := 2348
	result, err := intToRoman(inputInt)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println("Result is ", result)
}

func intToRoman(inputInt int) (string, error) {
	if inputInt > 4000 {
		return "", errors.New("invalid number")
	}

	return convertIntToRoman(inputInt), nil
}

func convertIntToRoman(inputInt int) string {
	var result string
	var division int
	var fullDivision int
	for i := 1000; i > 0; {
		division = inputInt / i
		fullDivision = division * i
		result = result + intRoman[fullDivision]
		inputInt = inputInt - fullDivision
		i = i / 10
	}

	return result
}
