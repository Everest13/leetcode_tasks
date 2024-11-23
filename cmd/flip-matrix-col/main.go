package main

import "fmt"

func main() {
	matrix := [][]int{{0,0,0},{0,0,1},{1,1,0}}
	//matrix = [][]int{}

	res := maxEqualRowsAfterFlips(matrix)
	fmt.Println(res)
}

/*
	1072. Flip Columns For Maximum Number of Equal Rows
	Medium
	https://leetcode.com/problems/flip-columns-for-maximum-number-of-equal-rows/?envType=daily-question&envId=2024-11-22
	Runtime 157ms Beats 50.00%, Memory 9.34MB Beats 40.00%
 */
func maxEqualRowsAfterFlips2(matrix [][]int) int {
	mH := len(matrix)
	mW := len(matrix[0])

	patternMap := map[string]int{}
	countIdentical := 0
	for i:=0; i<mH; i++ {
		identical := true
		s := "*"
		for j:=1; j<mW; j++ {
			if matrix[i][j] == matrix[i][j-1] {
				s += "*"
				continue
			}
			s += "|*"
			identical = false
		}
		patternMap[s]++
		if identical {
			countIdentical++
		}
	}

	maxCount := countIdentical
	for _, count := range patternMap {
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}

/*
	Runtime 4ms Beats 75.00%, Memory 11.18MB Beats 20.00%
 */
func maxEqualRowsAfterFlips(matrix [][]int) int {
	mH := len(matrix)
	mW := len(matrix[0])

	patternMap := make(map[string]int)
	maxCount := 0

	for i := 0; i < mH; i++ {
		// Form key that will be the same for lines with the same patterns after reversal
		key := make([]byte, mW)
		for j := 0; j < mW; j++ {
			if matrix[i][j] == matrix[i][0] {
				key[j] = '0' // '0' for matches
			} else {
				key[j] = '1' // '1' for differences
			}
		}
		keyStr := string(key)
		patternMap[keyStr]++
		if patternMap[keyStr] > maxCount {
			maxCount = patternMap[keyStr]
		}
	}

	return maxCount
}
