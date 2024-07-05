package main

import "fmt"

func main() {
	height := []int{0,1,0,2,1,0,1,3,2,1,2,1} //6

	height = []int{4,2,0,3,2,5} //9

	res := trap(height)
	fmt.Println(res)
}

/*
	42. Trapping Rain Water
	Hard
	https://leetcode.com/problems/trapping-rain-water/?envType=daily-question&envId=2024-05-11
	Runtime 3ms Beats 98.78%, Memory 6.08MB Beats 36.47%
 */
func trap(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}

	maxEl := height[0]
	maxElId := 0
	for i, val := range height {
		if val > maxEl {
			maxEl = val
			maxElId = i
		}
	}

	water := 0
	maxBefore := height[0]
	for i:=1; i<maxElId; i++ {
		h := height[i]
		if h < maxBefore {
			water += maxBefore - h
			continue
		}
		maxBefore = h
	}

	maxBefore = height[n-1]
	for i:=n-2; i>maxElId; i-- {
		h := height[i]
		if h < maxBefore {
			water += maxBefore - h
			continue
		}
		maxBefore = h
	}

	return water
}