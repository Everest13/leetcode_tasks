package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

/*
	1980. Find Unique Binary String
	https://leetcode.com/problems/find-unique-binary-string/description/?envType=daily-question&envId=2023-11-16
	Runtime 0 ms Beats 100%, Memory 2.3 MB Beats 56.26% M
*/
func main() {
	test := []string{"00","10"}
	test = []string{"00001101111011","11011001010101","00001000000001","00011100000000","01101111001010","01011011010101",
		"00110100111111","10001100101110","01000000010011","00110000101011","01101010000101","10000111010100","11000000111000",
		"00101101110110"}
	test = []string{"0"}
	//test = []string{"0000000110","0000000011","0000000001","0000000000","0000001000","0000000100","0000000101","1111111111","0000000111","0000001001"}

	res := findDifferentBinaryString(test)
	fmt.Println("Res is", res)
}

func findDifferentBinaryString(nums []string) string {
	var intNums []int
	for _, num := range nums {
		intNum, _ := strconv.Atoi(num)
		intNums = append(intNums, intNum)
	}
	sort.Ints(intNums)

	n :=len(nums)
	combs := int(math.Pow(2, float64(n)))
	baseNums := []int{0,1}

	convert := func(num int) string {
		str := strconv.FormatInt(int64(num), 10)
		zero := strings.Repeat("0", n-len(str))
		return zero+str
	}

	for i:=0; i < 2; i++ {
		if i == n || baseNums[i] != intNums[i] {
			return convert(baseNums[i])
		}
	}

	b := 10
	for i:=2; i < combs; {
		for _, bN := range baseNums {
			if i == n || b+bN != intNums[i] {
				return convert(b+bN)
			}

			baseNums = append(baseNums, b+bN)
			i++
		}
		b*=10
	}

	return ""
}

