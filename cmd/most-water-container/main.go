package main

import (
	"fmt"
)

func main() {
	height := []int{1,8,6,2,5,4,8,3,7} //49

	res := maxArea(height)
	fmt.Println(res)
}
/*
	11. Container With Most Water
	Medium
	https://leetcode.com/problems/container-with-most-water/
	Runtime 60ms Beats 72.84%, Memory 7.68MB Beats 76.21%
 */
func maxArea(height []int) int {
	getSquare := func(i, j int) int {
		p := j-i
		if height[i] < height[j] {
			return p * height[i]
		}

		return  p * height[j]
	}

	i := 0
	j := len(height)-1
	maxSquare := getSquare(i, j)
	for i < j {
		if height[i] > height[j] {
			j--
			if height[j] < height[j+1] {
				continue
			}
		} else {
			i++
			if height[i] < height[i-1] {
				continue
			}
		}

		square := getSquare(i, j)
		if square > maxSquare {
			maxSquare = square
		}
	}

	return maxSquare
}



// Time Limit Exceeded
func maxArea2(height []int) int {
	getSquare := func(i, j int) int {
		p := j-i
		if height[i] < height[j] {
			return p * height[i]
		}

		return  p * height[j]
	}

	n := len(height)
	maxSquare := 0
	for i:=0; i<n; i++ {
		for j:=1; j<n; j++ {
			square := getSquare(i,j)
			if square > maxSquare {
				maxSquare = square
			}
		}
	}

	return maxSquare
}
