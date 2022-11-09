package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const regexRoman = "^[I,V,X,L,C,D,M]{0,15}$"

var romanInt = map[string]int{
	"M": 1000,
	"D": 500,
	"C": 100,
	"L": 50,
	"X": 10,
	"V": 5,
	"I": 1,
}

func main() {
	// input roman like "LVIII", "MMMCCCXCIX", "MMCDXXXV"
	inputRoman := "MMCDXXXV"
	result, err := romanToInt(inputRoman)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println("Result is ", result)
}

// Task #1
// Given a roman numeral, convert it to an integer.
// https://leetcode.com/problems/roman-to-integer/
// Runtime: 19 ms, faster than 33.90% of Go online submissions for Roman to Integer.
// Memory Usage: 2.9 MB, less than 84.26% of Go online submissions for Roman to Integer.
func romanToInt(romanNumber string) (int, error) {
	validRomanNumber, err := validateRoman(romanNumber)
	if err != nil {
		return 0, err
	}

	return convertRomanToInt(validRomanNumber), nil
}

func convertRomanToInt(romanNumber string) int {
	var result int
	var currentVal int
	var nextVal int
	lenChars := len(romanNumber)
	for i := 0; i < lenChars; i++ {
		currentVal = romanInt[string(romanNumber[i])] //получаем значение текущего символа
		j := i + 1
		if j < lenChars { //если символ последний - не заходим на проверку
			nextVal = romanInt[string(romanNumber[j])] //получаем следующий символ
			if nextVal > currentVal {
				result += nextVal - currentVal //если больше - высчитаваем исключение
				i = j
				continue
			}
		}
		result += currentVal
	}

	return result
}

func validateRoman(romanNumber string) (string, error) {
	strings.TrimSpace(romanNumber)
	r, err := regexp.Compile(regexRoman)
	if err != nil {
		return "", err
	}

	result := r.MatchString(romanNumber)
	if !result {
		return "", errors.New("invalid roman number")
	}

	return romanNumber, nil
}
