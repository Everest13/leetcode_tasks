package main

import "fmt"

/*
	6. Zigzag Conversion
	https://leetcode.com/problems/zigzag-conversion/description/
	Runtime - 3ms(beats 94.7%), memory: 5.5mb(beats 73.1%) M
*/
func main() {
	str := "PAYPALISHIRING"
	str = "AB"
	numRows := 1

	result := convert(str, numRows)
	fmt.Println(result)
}

func convert(s string, numRows int) string {
	var result []rune
	sArr := []rune(s)
	lenSArr := len(sArr)
	zzSymbs := 0 //zigzag

	if numRows > 2 {
		zzSymbs = numRows - 2
	}

	for i := 0; i < numRows; i++ {
		j := i
		if i == 0 || i == numRows-1 {
			for j < lenSArr {
				result = append(result, sArr[j])
				j = j + numRows + zzSymbs
			}
			continue
		}

		for j < lenSArr {
			result = append(result, sArr[j])
			j = j + numRows + zzSymbs
			k := j - 2*i
			if k < lenSArr {
				result = append(result, sArr[k])
			}
		}
	}

	return string(result)
}
