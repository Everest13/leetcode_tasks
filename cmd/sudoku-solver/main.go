package main

import "fmt"

func main() {
	//в каждой строке 1-9 без повторений
	//в каждой колонке 1-9 без повторений
	// в каждом боксе  1-9 без повторений
	board := [][]byte{
		{'5','3','.','.','7','.','.','.','.'},
		{'6','.','.','1','9','5','.','.','.'},
		{'.','9','8','.','.','.','.','6','.'},
		{'8','.','.','.','6','.','.','.','3'},
		{'4','.','.','8','.','3','.','.','1'},
		{'7','.','.','.','2','.','.','.','6'},
		{'.','6','.','.','.','.','2','8','.'},
		{'.','.','.','4','1','9','.','.','5'},
		{'.','.','.','.','8','.','.','7','9'}}

	solveSudoku(board)
}

type key struct {
	line int
	col int
	boxN int
}

func solveSudoku(board [][]byte)  {

	lineAbsentMap := map[int]map[byte]struct{}{} //пропуски по линии
	columnAbsentMap := map[int]map[byte]struct{}{} //пропуски по колонне
	boxAbsentMap := map[int]map[byte]struct{}{} //пропуски по боксу
	absentVals := map[key][]byte{} //все пропуски
	absentVals2 := map[*[][]byte][]byte{} //все пропуски

	//попробовать copy с base
	//todo попробовать реализацию через поинтеры
	getAbsentVals := func(line, col int) {
		lineMap, ok1 := lineAbsentMap[line]
		colMap, ok2 := columnAbsentMap[col]

		idX, idY := 0, 0
		for i:=0; i<9; i++ {
			baseLine := map[byte]struct{} {'1':{}, '2':{}, '3':{}, '4':{}, '5':{}, '6':{}, '7':{}, '8':{}, '9':{}}
			baseColumn := map[byte]struct{} {'1':{}, '2':{}, '3':{}, '4':{}, '5':{}, '6':{}, '7':{}, '8':{}, '9':{}}
			baseBox := map[byte]struct{} {'1':{}, '2':{}, '3':{}, '4':{}, '5':{}, '6':{}, '7':{}, '8':{}, '9':{}}

			if board[line][i] != '.' {
				delete(baseLine, board[line][i])
			}

			if board[i][col] != '.' {
				delete(baseColumn, board[i][col])
			}

			//box
			for m := 0; m < 3; m++ {
				for k := 0; k < 3; k++ {
					if board[idX][idY] != '.' {
						delete(baseBox, board[idX][idY])
					} else {
						absentVals[key{line: m, col: k, boxN: i}] = []byte{}
					}
					idY++
				}
				idY = idY - 3
				idX++
			}
			if idX >= 9 {
				idX = 0
				idY = idY + 3
			}

			lineAbsentMap[i] = baseLine
			columnAbsentMap[i] = baseColumn
			boxAbsentMap[i] = baseBox
		}
	}

	//набьираем возможные значения

	for i, line := range board {
		for j, colm := range line {





		}
	}







	// собрать возможные значения по строкам, колонкам, боксам

	for i:=0; i<9; i++ {
		baseLine := map[byte]struct{} {'1':{}, '2':{}, '3':{}, '4':{}, '5':{}, '6':{}, '7':{}, '8':{}, '9':{}}
		baseColumn := map[byte]struct{} {'1':{}, '2':{}, '3':{}, '4':{}, '5':{}, '6':{}, '7':{}, '8':{}, '9':{}}
		baseBox := map[byte]struct{} {'1':{}, '2':{}, '3':{}, '4':{}, '5':{}, '6':{}, '7':{}, '8':{}, '9':{}}

		for j:=0; j<9; j++ {
			if board[i][j] != '.' {
				delete(baseLine, board[i][j])
			}

			if board[j][i] != '.' {
				delete(baseColumn, board[j][i])
			}
		}

		//box
		for m := 0; m < 3; m++ {
			for k := 0; k < 3; k++ {
				if board[idX][idY] != '.' {
					delete(baseBox, board[idX][idY])
				} else {
					absentVals[key{line: m, col: k, boxN: i}] = []byte{}
				}
				idY++
			}
			idY = idY - 3
			idX++
		}
		if idX >= 9 {
			idX = 0
			idY = idY + 3
		}

		lineAbsentMap[i] = baseLine
		columnAbsentMap[i] = baseColumn
		boxAbsentMap[i] = baseBox
	}

	//absentVals := map[key][]byte{} //все пропуски
	for len(absentVals) > 0 {

		for k := range absentVals {
			if len(absentVals[k]) == 1 {
				board[k.line][k.col] = absentVals[k][0]
				delete(absentVals, k)
			}





		}
	}



	fmt.Println(board)
}
