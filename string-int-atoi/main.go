package main

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var regex = `^\s+([+-]?\d+)`

/*
8. String to Integer (atoi)
https://leetcode.com/problems/string-to-integer-atoi/submissions/
Runtime - 4ms(beats 43.84%), memory: 2.3mb(beats 29.2%)
*/
func main() {
	test := "   -00042"
	//test = "20000000000000000000"
	//test = " +-"
	//test = "      -4234324fvfdvf-23232"
	//test = "9223372036854775808"

	result := myAtoi(test)
	fmt.Println(result)
}

func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")

	minus := false
	if len(s) > 1 {
		if s[0] == '-' {
			minus = true
			s = s[1:]
		} else if s[0] == '+' {
			s = s[1:]
		}
	}

	s = strings.TrimLeft(s, "0")

	getInt := func(char rune) (int, error) {
		switch char {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return int(char - '0'), nil
		}
		return 0, errors.New("")
	}

	result := 0
	for _, char := range s {
		number, err := getInt(char)
		if err != nil {
			break
		}

		if result == 0 {
			result = number
			continue
		}

		result = result*10 + number

		if result > math.MaxInt32 {
			if minus {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}

	if minus {
		result = 0 - result
	}

	return result
}

/*
Use regex
Runtime - 4ms(beats 43.84%), memory: 6mb(beats 6.4%)
*/
func myAtoi2(s string) int {
	re := regexp.MustCompile(regex)
	newSArr := re.FindStringSubmatch(s)

	var newS string
	if len(newSArr) > 1 {
		newS = newSArr[1]
	} else {
		return 0
	}

	minus := false
	if newS[0] == '-' {
		minus = true
		newS = newS[1:]
	}
	if len(newS) == 0 {
		return 0
	}

	newS = strings.TrimLeft(newS, "0")
	if len(newS) == 0 {
		return 0
	}

	result, err := strconv.Atoi(newS)
	if err != nil {
		if minus {
			return math.MinInt32
		}
		return math.MaxInt32
	}

	if minus {
		result = 0 - result
	}

	if result > math.MaxInt32 {
		return math.MaxInt32
	}

	if result < math.MinInt32 {
		return math.MinInt32
	}

	return result
}
