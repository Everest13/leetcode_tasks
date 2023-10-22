package main

import "fmt"

/*
	36. Valid Sudoku. Medium.
	https://leetcode.com/problems/valid-sudoku/description/
	Runtime 1 ms Beats 61.77%, Memory 3.3 MB Beats 40.61%
 */
func main() {
	board := [][]byte{
		 //в каждой строке 1-9 без повторений
		//в каждой колонке 1-9 без повторений
		// в каждом боксе  1-9 без повторений

		 //1 //2 //3 //4 //5 //6 //7 //8 //9
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
	mapLine := map[byte]struct{}{}
	mapBoard := map[byte]struct{}{}
	idX, idY := 0, 0
	for i:=0; i<9; i++ {
		for j:=0; j<9; j++ {
			val1 := board[i][j] //line
			val2 := board[j][i] //column
			if val1 != '.' {
				if _, ok := mapLine[val1]; ok {
					return false
				}

				mapLine[val1] = struct{}{}
			}

			if val2 != '.' {
				if _, ok := mapBoard[val2]; ok {
					return false
				}

				mapBoard[val2] = struct{}{}
			}
		}
		mapLine = map[byte]struct{}{}
		mapBoard = map[byte]struct{}{}

		//box
		for m := 0; m < 3; m++ {
			for k := 0; k < 3; k++ {
				val := board[idX][idY]
				idY++

				if val != '.' {
					if _, ok := mapBoard[val]; ok {
						return false
					}
					mapBoard[val] = struct{}{}
				}
			}
			idY = idY - 3
			idX++
		}

		mapBoard = map[byte]struct{}{}
		if idX >= 9 {
			idX = 0
			idY = idY + 3
		}

		mapBoard = map[byte]struct{}{}
	}
	return true
}