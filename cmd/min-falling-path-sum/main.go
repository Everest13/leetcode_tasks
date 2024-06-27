package main

import "fmt"

func main() {
	matrix := [][]int{{2,1,3},{6,5,4},{7,8,9}} //13

	matrix = [][]int{{17,82},{1,-44}} //-27

	res := minFallingPathSum(matrix)
	fmt.Println(res)
}

/*
	931. Minimum Falling Path Sum
	Medium
	https://leetcode.com/problems/minimum-falling-path-sum/?envType=daily-question&envId=2024-05-11
	Runtime 7ms Beats 87.59%, Memory 4.85MB Beats 77.37%
*/
func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)

	for r:=1; r<n; r++ {
		rUp := r-1
		for c:=0; c<n; c++ {
			minAbove := matrix[rUp][c]
			if c>=1 && matrix[rUp][c-1] < minAbove {
				minAbove = matrix[rUp][c-1]
			}
			if c+1<n && matrix[rUp][c+1] < minAbove {
				minAbove = matrix[rUp][c+1]
			}

			matrix[r][c] += minAbove
		}
	}

	n--
	minSumm := matrix[n][0]
	for c:=1; c<=n; c++ {
		if matrix[n][c] < minSumm {
			minSumm = matrix[n][c]
		}
	}

	return minSumm
}
