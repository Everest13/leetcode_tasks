package main

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
	boxN int
}

func solveSudoku(board [][]byte) { // strconv.Atoi(string(val))

	existLines := map[int]map[byte]struct{}{}
	existCols := map[int]map[byte]struct{}{}
	existBoxes := map[int]map[byte]struct{}{}

	possibleVals := map[key]map[byte]struct{}{}

	for i := 0; i < 9; i++ {
		for j := i; j < 9; j++ {
			if board[i][j] == '.' {
				boxN1 := checkBoxN(i, j)
				possibleVals[key{x:i, y:j, boxN: boxN1}] = map[byte]struct{}{'1': {}, '2': {}, '3': {}, '4': {}, '5': {}, '6': {}, '7': {}, '8': {}, '9': {}}
			}
			if i == j {
				continue
			}
			if board[j][i] == '.' {
				boxN2 := checkBoxN(j, i)
				possibleVals[key{x:j, y:i, boxN: boxN2}] = map[byte]struct{}{'1': {}, '2': {}, '3': {}, '4': {}, '5': {}, '6': {}, '7': {}, '8': {}, '9': {}}
			}
		}
	}

	cleanPossibleVals := func(i, j int, val byte) {
		boxN1 := checkBoxN(i, j)
		for k := range possibleVals {
			if k.x == i || k.y == j || k.boxN == boxN1 {
				delete(possibleVals[k], val)
			}
			if i == j {
				continue
			}
			boxN1 = checkBoxN(j, i)
			if k.x == j || k.y == i || k.boxN == boxN1 {
				delete(possibleVals[k], val)
			}
		}
	}

	for k, vals := range possibleVals {
		if _, ok := existLines[k.x]; ok {
			delete(possibleVals[k], vals)
		}

	}





	return
}


func checkBoxN(i, j int) int {
	switch {
	case i<3 && j<3: //1
		return 1
	case i<6 && j<3: //2
		return 2
	case i<9 && j<3: //3
		return 3
	case i<3 && j<6: //4
		return 4
	case i<6 && j<6: //5
		return 5
	case i<9 && j<6: //6
		return 6
	case i<3 && j<9: //7
		return 7
	case i<6 && j<9: //8
		return 8
	default: // i<9 && j<9: //9
		return 9
	}
}

func solveSudoku2(board [][]byte) {
	existLines := map[int]map[byte]struct{}{}
	existCols := map[int]map[byte]struct{}{}
	existBoxes := map[int]map[byte]struct{}{}
	absentValBoxes := map[int]map[byte]struct{}{}
	possibleCeilVals := map[key][]byte{}

	idX, idY := 0, 0
	for i := 0; i < 9; i++ {
		line := map[byte]struct{}{}
		col := map[byte]struct{}{}

		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				line[board[i][j]] = struct{}{}
			}

			if board[j][i] != '.' {
				col[board[j][i]] = struct{}{}
			}
		}
		existLines[i] = line
		existCols[i] = col

		existBox := map[byte]struct{}{} //сущ значения в боксе
		// возможные значения для бокса
		absentValBox := map[byte]struct{}{'1': {}, '2': {}, '3': {}, '4': {}, '5': {}, '6': {}, '7': {}, '8': {}, '9': {}}
		m := idX + 3
		n := idY + 3
		for idX < m {
			for idY < n {
				val := board[idY][idX]
				idY++
				if val != '.' {
					existBox[val] = struct{}{}
					delete(absentValBox, val)
					continue
				}
				//
				possibleCeilVals[key{x:idX, y:idY, boxN: i}] = []byte{}
			}
			idX++
			if idX == 9 {
				idX = 0
				break
			}
			idY -= 3
		}

		existBoxes[i] = existBox
		absentValBoxes[i] = absentValBox
	}

	checkVal := func(k key, val byte) bool {
		_, ok1 := existLines[k.x][val]
		_, ok2 := existCols[k.y][val]
		_, ok3 := existBoxes[k.boxN][val]

		if !ok1 && !ok2 && !ok3 {
			return true
		}

		return false
	}

	addVal := func(k key, val byte)  {
		existLines[k.x][val] = struct{}{}
		existCols[k.y][val] = struct{}{}
		existBoxes[k.boxN][val] = struct{}{}
	}

	track := []map[key]byte{} //добавленные значения
	for k, vals := range possibleCeilVals {

		if len(vals) == 0 { //найти значения если их нет
			for val := range absentValBoxes[k.x] {
				_, okLine := existLines[k.x][val]
				_, okCol := existCols[k.y][val]
				if !okLine && !okCol {
					vals = append(vals, val)
				}
			}
		}

		for i, val := range vals {
			if checkVal(k, val) {
				vals = vals[i:] //удалить значения которые не подходят  todo? i
				track = append(track, map[key]byte{k:val})
				addVal(k, val)
				if len(vals) == 0 {
					delete(possibleCeilVals, k)
				}

				possibleCeilVals[k] = vals
			}


		}





	}



	return
}