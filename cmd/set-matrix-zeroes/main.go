package main

import "fmt"

func main() {
	matrix := [][]int{{1,1,1},{1,0,1},{1,1,1}}
	matrix = [][]int{{1,0,2,0},{3,4,5,2},{1,3,1,5}}

	setZeroes(matrix)

	fmt.Println(matrix)
}

/*
	73. Set Matrix Zeroes
	Medium
	https://leetcode.com/problems/set-matrix-zeroes/description/
	Runtime 0ms Beats 100.00%,, Memory 7.52MB Beats 8.61%
 */
func setZeroes(matrix [][]int)  {
	zeroArr := [][2]int{}
	for r:=0; r<len(matrix); r++ {
		for c:=0; c<len(matrix[0]); c++ {
			if matrix[r][c] == 0 {
				zeroArr = append(zeroArr, [2]int{r,c})
			}
		}
	}

	for _, arr := range zeroArr {
		r := arr[0]
		c := arr[1]
		for i:=0; i<len(matrix[0]); i++ {
			matrix[r][i] = 0
		}

		for i:=0; i<len(matrix); i++ {
			matrix[i][c] = 0
		}
	}
}