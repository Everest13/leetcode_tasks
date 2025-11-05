package main

import "fmt"

/*
	36. Valid Sudoku
	Medium
	https://leetcode.com/problems/valid-sudoku/description/
 */
func main() {
	board := [][]byte{
		{'5','3','.','.','7','.','.','.','.'}, //1
		{'6','.','.','1','9','5','.','.','.'}, //2
		{'.','9','8','.','.','.','.','6','.'}, //3
		{'8','.','.','.','6','.','.','.','3'}, //4
		{'4','.','.','8','.','3','.','.','1'}, //5
		{'7','.','.','.','2','.','.','.','6'}, //6
		{'.','6','.','.','.','.','2','8','.'}, //7
		{'.','.','.','4','1','9','.','.','5'}, //8
		{'.','.','.','.','8','.','.','7','9'}, //9
	}

	res := isValidSudoku(board)
	fmt.Println("Result is ", res)
}
func isValidSudoku(board [][]byte) bool {
	rows := [9][9]bool{}
	cols :=  [9][9]bool{}
	boxes := [9][9]bool{}

	for i:=0; i<9; i++ {
		boardRow := board[i]
		for j:=0; j<9; j++ {
			if boardRow[j] == '.' {
				continue
			}

			boardIJ := int(boardRow[j] - '0')-1
			n := boxDetermine(i,j)

			if rows[i][boardIJ] || cols[j][boardIJ] || boxes[n][boardIJ] {
				return false
			}

			rows[i][boardIJ] = true
			cols[j][boardIJ] = true
			boxes[n][boardIJ] = true
		}
	}

	return true
}

func boxDetermine(i,j int) int {
	return (i/3)*3 + (j/3)
}
