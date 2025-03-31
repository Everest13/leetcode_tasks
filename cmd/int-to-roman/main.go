package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 3749
	//num = 2049
	num = 1994

	res := intToRoman(num)
	fmt.Println(res)
}

var romansMap = map[int]string {1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX",
	10: "X", 20: "XX", 30: "XXX", 40: "XL", 50: "L", 60: "LX", 70: "LXX", 80: "LXXX", 90: "XC",
	100: "C", 200: "CC", 300: "CCC", 400: "CD", 500: "D", 600: "DC", 700: "DCC", 800: "DCCC",
	900: "CM", 1000: "M", 2000: "MM", 3000: "MMM",
}

/*
	12. Integer to Roman
	Medium
	https://leetcode.com/problems/integer-to-roman/description/?envType=study-plan-v2&envId=top-interview-150
	Runtime 0ms Beats 100.00%, Memory 5.21MB Beats 29.42%
 */
func intToRoman(num int) string {
	res := ""
	digitNum := len(strconv.Itoa(num))
	divided := 1
	switch digitNum {
	case 4:
		divided *= 1000
	case 3:
		divided *= 100
	case 2:
		divided *= 10
	}

	for divided > 0 {
		quotient := num/divided*divided
		res += romansMap[quotient]
		num -= quotient
		divided /= 10
	}

	return res
}

var romaMap = map[int]string {1: "I", 5: "V", 10: "X", 50: "L", 100: "C", 500: "D", 1000: "M"}

/*
	Runtime 0ms Beats 100.00%, Memory 5.12MB Beats 50.96%
 */
func intToRoman2(num int) string {
	res := ""
	convert := func(quotn, divid int) string {
		switch quotn {
		case 1:
			return romaMap[quotn*divid]
		case 2:
			roman1 := romaMap[1*divid]
			return roman1+roman1
		case 3:
			roman1 := romaMap[1*divid]
			return roman1+roman1+roman1
		case 4:
			return romaMap[1*divid]+romaMap[5*divid]
		case 5:
			return romaMap[5*divid]
		case 6:
			return romaMap[5*divid]+romaMap[1*divid]
		case 7:
			roman1 := romaMap[1*divid]
			return romaMap[5*divid]+roman1+roman1
		case 8:
			roman1 := romaMap[1*divid]
			return romaMap[5*divid]+roman1+roman1+roman1
		case 9:
			return romaMap[1*divid]+romaMap[10*divid]
		case 10:
			return romaMap[10*divid]
		}
		return ""
	}

	numStr := strconv.Itoa(num)
	digitNum := len(numStr)
	divided := 1
	switch digitNum {
	case 4:
		divided *= 1000
	case 3:
		divided *= 100
	case 2:
		divided *= 10
	}

	for _, char := range numStr {
		digit := int(char - '0')
		res += convert(digit, divided)
		divided /= 10
	}

	return res
}

