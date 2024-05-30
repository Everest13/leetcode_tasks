package main

import (
	"errors"
	"fmt"
)

var numberChars = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

/*
	17. Letter Combinations of a Phone Number
	https://leetcode.com/problems/letter-combinations-of-a-phone-number/
	Runtime: 0 ms, faster than 100.00%, Memory Usage: 2.1 MB, less than 16.39% M
 */
func main() {
	digits := "234"
	err := validation(digits)
	if err != nil {
		fmt.Println("Error is ", err.Error())
		return
	}

	result := letterCombinations(digits)
	fmt.Println("Result is ", result)
}

func validation(digits string) error {
	if len(digits) == 0 || len(digits) > 4 {
		return errors.New("impossible letter combinations")
	}

	return nil
}

func letterCombinations(digits string) []string {
	var result []string

	sets := make([][]string, 0, len(digits))
	for _, char := range digits {
		sets = append(sets, numberChars[string(char)])
	}

	for _, set := range sets {
		if len(result) == 0 {
			result = set
			continue
		}

		newComb := make([]string, 0, len(set)*len(result))
		for _, char := range set {
			for _, combination := range result {
				newComb = append(newComb, combination+char)
			}
		}
		result = newComb
	}

	return result
}
