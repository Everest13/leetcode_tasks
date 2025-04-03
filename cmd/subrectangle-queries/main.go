package main

func main() {

}

/**
	1476. Subrectangle Queries
	Medium
	https://leetcode.com/problems/subrectangle-queries/description/
 */
type SubrectangleQueries struct {
	rectangle [][]int
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int)  {
	for i:=row1; i<=row2; i++ {
		for j:=col1; j<=col2; j++ {
			this.rectangle[i][j] = newValue
		}
	}
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	return this.rectangle[row][col]
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{rectangle}
}




