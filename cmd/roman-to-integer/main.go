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
		fmt.Print(err.Error())
		return
	}

	fmt.Println("Result is ", result)
}

/*
	13. Given a roman numeral, convert it to an integer.
	https://leetcode.com/problems/roman-to-integer/
	Runtime: 11 ms, faster than 74.44% Memory Usage: 3.5 MB, less than 15.10% E
 */
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
	romanArr := strings.Split(romanNumber, "")
	for i := 0; i < lenChars; i++ {
		currentVal = romanInt[romanArr[i]]
		j := i + 1
		if j < lenChars { //если символ последний - не заходим на проверку
			nextVal = romanInt[romanArr[j]] //получаем следующий символ
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
