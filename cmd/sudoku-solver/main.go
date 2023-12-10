package main


/*
	37. Sudoku Solver
	Hard
	https://leetcode.com/problems/sudoku-solver/description/?envType=list&envId=r57h7kzc
	Runtime 90 ms Beats 5.51%, Memory 8.1 MB Beats 5.51%
 */
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
	x int
	y int
}

func solveSudoku(board [][]byte) {
	track := []key{}
	unapproprVals := map[key]map[byte]struct{}{}

	addVal := func(k key, absentVals map[byte]struct{} ) {
		for val := range absentVals {
			track = append(track, k)
			board[k.x][k.y] = val
			break
		}
	}

	rollBack := func(i, j int) (int, int) {

		lenTrack := len(track) - 1
		k := track[lenTrack]
		track = track[:lenTrack]

		delete(unapproprVals, key{x:i, y:j})

		val := board[k.x][k.y]
		board[k.x][k.y] = '.'
		if _, ok := unapproprVals[k]; ok {
			unapproprVals[k][val] = struct{}{}
			return k.x, k.y
		}

		unapproprVals[k] = map[byte]struct{}{val: {}}

		return k.x, k.y
	}

	findAbsentVals := func(k key) map[byte]struct{} {
		base := map[byte]struct{}{'1': {}, '2': {}, '3': {}, '4': {}, '5': {}, '6': {}, '7': {}, '8': {}, '9': {}}
		sI := k.x/3 * 3
		sJ := k.y/3 * 3

		m, n := sI, sJ
		for i := 0; i < 9; i++ {
			val1 := board[k.x][i]
			if val1 != '.' {
				delete(base, val1)
			}

			val2 := board[i][k.y]
			if val2 != '.' && val1 != val2 {
				delete(base, val2)
			}

			//	box
			val3 := board[m][n]
			if val3 != '.' && val1 != val3 && val2 != val3 {
				delete(base, val3)
			}
			n++
			if n == (sJ+3) {
				m++
				n = sJ
			}
		}

		for val := range unapproprVals[k] {
			delete(base, val)
		}

		return base
	}

	//1
	for i := 0; i < 9; {
		for j := 0; j < 9; {
			valL := board[i][j]
			if valL == '.' {
				k := key{x:i, y:j}

				absentVals := findAbsentVals(k)
				if len(absentVals) == 0 {
					i,j = rollBack(i,j)
					continue
				}

				addVal(k, absentVals)
			}
			j++
		}
		i++
	}

	return
}
