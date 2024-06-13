package main

import "fmt"

func main() {
	s := "12" //2
	s = "226" //3

	s = "10" //1
	s = "2101" //1

	res := numDecodings(s)
	fmt.Println(res)
}

/*
	91. Decode Ways
	https://leetcode.com/problems/decode-ways/?envType=daily-question&envId=2024-05-11
 */
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	convertToInt := func(c byte) uint64 {
		var d byte
		d = c - '0'
		return uint64(d)
	}

	lenS := len(s)
	decodeCount := 0
	stopRec := false
	recursion := func(i int) {}
	recursion = func(i int) {
		if stopRec == true {
			return
		}

		for j:=i; j<lenS; j++ {

			isZero := s[j] == '0'
			nextZero := j+1<lenS && s[j+1] == '0'
			pastZero := s[j-1] == '0'

			if isZero {
				if nextZero || pastZero {
					stopRec = false
					decodeCount = 0
					return
				}
				j++
				continue
			}

			if nextZero {
				j+=2
				continue
			}

			num := convertToInt(s[j-1])*10 + convertToInt(s[j])
			if num <= 26 {
				recursion(j+2)
			}
		}

		decodeCount++
	}

	recursion(1)

	return  decodeCount
}