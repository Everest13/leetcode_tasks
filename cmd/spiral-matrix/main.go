package main

import "fmt"

/**
	54. Spiral Matrix. Medium.
	https://leetcode.com/problems/spiral-matrix/description/?envType=daily-question&envId=2023-10-05
	Runtime 1ms, beats 80,6%; Memory 1.9mb, beats 51.6%.
*/
func main() {

	matrix := [][]int{{1,2,3}, {4,5,6}, {7,8,9}}
	//matrix = [][]int{{1,2,3,4}, {5,6,7,8}, {9,10,11,12}}
	matrix = [][]int{{1,2,3,4,5,6}, {7,8,9,10,11,12}, {13,14,15,16,17,18}, {19,20,21,22,23,24}, {25,26,27,28,29,30},
		{31,32,33,34,35,36}, {37,38,39,40,41,42}}

	res := spiralOrder(matrix)
	fmt.Println("Result is ", res)
}

func spiralOrder(matrix [][]int) []int {

	mWdth := len(matrix[0]) //ширина
	mHht := len(matrix) //длина
	ln := mWdth*mHht
	res := make([]int, 0, ln)

	i, j, btm, lft := 0, 0, 0, 0
	mWdth--
	mHht--
	for len(res) < ln {
		//по порядку влево
		if len(res) < ln {
			for j <= mWdth {
				res = append(res, matrix[i][j])
				j++
			}

			i++
			j--
			lft++
		}

		// верт вниз вправо
		if len(res) < ln {
			for i <= mHht {
				res = append(res, matrix[i][j])
				i++
			}

			j--
			i--
			mWdth--
		}

		//обратный порядок
		if len(res) < ln {
			for j >= btm {
				res = append(res, matrix[i][j])
				j--
			}

			j++
			i--
			mHht--
		}

		//верт вверх лево
		if len(res) < ln {
			for i >= lft {
				res = append(res, matrix[i][j])
				i--
			}

			j++
			i++
			btm++
		}
	}

	return res
}
